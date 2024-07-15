package repository

import (
	"database/sql"
	"st/backend/db"
	"st/backend/db/entity"
)

type MetadataRepository struct {
	db        *sql.DB
	modelName string
}

var _ db.Repository = (*MetadataRepository)(nil)

// func (r *MetadataRepository) FindAll() []entity.Metadata {
// 	metadatas := make()
// }

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
	return migrate(r.db, r.modelName, entity.MetadataTableVersion, &entity.User{})
}
