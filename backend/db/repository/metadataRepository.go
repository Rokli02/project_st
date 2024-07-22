package repository

import (
	"database/sql"
	"fmt"
	"st/backend/db"
	"st/backend/db/entity"
	"st/backend/utils"
	"st/backend/utils/logger"
	"st/backend/utils/settings"
	"strings"
	"time"
)

type MetadataRepository struct {
	db        *sql.DB
	modelName string
}

var _ db.Repository = (*MetadataRepository)(nil)

func (r *MetadataRepository) FindById(id int64) *entity.Metadata {
	row := r.db.QueryRow(fmt.Sprintf("SELECT key, value, type, updated_at, expire_at FROM %s WHERE id = ? LIMIT 1;", r.modelName), id)

	metadata := &entity.Metadata{Id: id}

	err := row.Scan(&(metadata.Key), &(metadata.Value), &(metadata.Type), &(metadata.UpdatedAt), &(metadata.ExpireAt))
	if err != nil {
		logger.WarningF("Couldn't find metadata in db for id '%d' (%s)", id, err)

		return nil
	}

	if isDateExpired(metadata.ExpireAt) {
		r.DeleteOne(id)

		metadata.Value = nil
		metadata.ExpireAt = nil
	}

	return metadata
}

func (r *MetadataRepository) FindAll() []entity.Metadata {
	var metadatas []entity.Metadata = make([]entity.Metadata, 0)
	var metadatasToRemove []int64 = make([]int64, 0)

	rows, _ := r.db.Query(fmt.Sprintf("SELECT id, key, value, type, updated_at, expire_at FROM %s;", r.modelName))
	for rows.Next() {
		metadata := entity.Metadata{}

		rows.Scan(&(metadata.Id), &(metadata.Key), &(metadata.Value), &(metadata.Type), &(metadata.UpdatedAt), &(metadata.ExpireAt))

		if isDateExpired(metadata.ExpireAt) {
			metadatasToRemove = append(metadatasToRemove, metadata.Id)

			metadata.Value = nil
			metadata.ExpireAt = nil
		}

		metadatas = append(metadatas, metadata)
	}

	r.DeleteMultiple(metadatasToRemove)

	return metadatas
}

func (r *MetadataRepository) InsertMultiple(metadatas []entity.Metadata) (inserted int) {
	stmt, _ := r.db.Prepare(fmt.Sprintf("INSERT INTO %s (key, value, type, updated_at, expire_at) VALUES (?,?,?,?,?);", r.modelName))
	defer stmt.Close()

	var updatedAt string = time.Now().Format(settings.Database.DateFormat)

	for _, metadata := range metadatas {
		var mutationArgs []any = []any{
			metadata.Key,      // Key
			metadata.Value,    // Value
			"app",             // Type
			updatedAt,         // UpdatedAt
			metadata.ExpireAt, // ExpireAt
		}

		if metadata.Type != "" {
			mutationArgs[2] = metadata.Type
		}

		res, err := stmt.Exec(mutationArgs...)
		if err != nil {
			logger.Warning("Some error occured during inserting metadata to db:", err)

			continue
		}

		affacted, _ := res.RowsAffected()
		inserted += int(affacted)
	}

	return
}

func (r *MetadataRepository) UpdateOne(id int64, metadata entity.Metadata) bool {
	var updatedAt string = time.Now().Format(settings.Database.DateFormat)
	template := fmt.Sprintf("UPDATE %s SET value = ?, type = ?, updated_at = %s, expire_at = ? WHERE id = ?;", r.modelName, updatedAt)

	var mutationArgs []any = []any{
		metadata.Value,    // Value
		"app",             // Type
		metadata.ExpireAt, // ExpireAt
		id,                // ID
	}

	if metadata.Type != "" {
		mutationArgs[1] = metadata.Type
	}

	res, err := r.db.Exec(template, mutationArgs...)
	affected, _ := res.RowsAffected()

	if err != nil || affected == 0 {
		logger.WarningF("Couldn't update %s record with id '%d' (%s)", r.modelName, id, err)

		return false
	}

	return true
}

func (r *MetadataRepository) DeleteOne(id int64) bool {
	// template := fmt.Sprintf("DELETE FROM %s WHERE id = ?;", r.modelName)
	template := fmt.Sprintf("UPDATE %s SET value = NULL, expire_at = NULL WHERE id = ?;", r.modelName)

	res, err := r.db.Exec(template, id)
	if err != nil {
		logger.ErrorF("Couldn't delete metadata record with id: %d (%s)", id, err)

		return false
	}
	affected, err := res.RowsAffected()
	if err != nil || affected == 0 {
		logger.WarningF("Couldn't delete record from %s table with id %d", r.modelName, id)

		return false
	}

	return true
}

func (r *MetadataRepository) DeleteMultiple(ids []int64) (deleted int) {
	if len(ids) < 1 {
		return
	}

	joinedIdsBuilder := strings.Builder{}
	for _, id := range ids {
		joinedIdsBuilder.WriteString(fmt.Sprintf("%d, ", id))
	}
	joinedIds := joinedIdsBuilder.String()

	// template := fmt.Sprintf("DELETE FROM %s WHERE id IN(%s);", r.modelName, joinedIds[:len(joinedIds)-2])
	template := fmt.Sprintf("UPDATE %s SET value = NULL, expire_at = NULL WHERE id IN(%s);", r.modelName, joinedIds[:len(joinedIds)-2])

	res, err := r.db.Exec(template)
	if err != nil {
		logger.ErrorF("Couldn't delete metadata record with ids: %v (%s)", ids, err)

		return 0
	}

	affected, err := res.RowsAffected()
	if err != nil || affected == 0 {
		logger.WarningF("Couldn't delete record from %s table with ids %v", r.modelName, ids)

		return 0
	}

	return int(affected)
}

func isDateExpired(expireDate *string) bool {
	return expireDate != nil && *expireDate != "" && utils.ToTime(*expireDate).Before(time.Now())
}

//
//	db.Repository Implementations
//

func (r *MetadataRepository) SetDB(db *sql.DB) {
	r.db = db
}

func (r *MetadataRepository) ModelName() string {
	return r.modelName
}

func (r *MetadataRepository) CreateTable() bool {
	return createTable(r.db, &entity.Metadata{}, r.modelName)
}

func (r *MetadataRepository) DropTable() bool {
	return dropTable(r.db, r.modelName)
}

func (r *MetadataRepository) IsTableExist() bool {
	return isTableExist(r.db, r.modelName)
}

func (r *MetadataRepository) Migrate() uint {
	return migrate(r.db, r.modelName, &entity.Metadata{})
}

func (r *MetadataRepository) InitTable() {
	userTableVersion := fmt.Sprintf("%d", (&entity.User{}).TableVersion())

	var initMetadatas []entity.Metadata = []entity.Metadata{
		{
			Key:   settings.MetadataKeys.UserTableVersion,
			Value: &userTableVersion,
		},
		{
			Key: settings.MetadataKeys.CurrentUserId,
		},
		{
			Key:   settings.MetadataKeys.LanguageId,
			Value: &settings.LanguageIds[0],
		},
	}

	all := len(initMetadatas)
	inserted := r.InsertMultiple(initMetadatas)

	logger.InfoF("Inserted into %s %d records out of %d", r.modelName, inserted, all)
}
