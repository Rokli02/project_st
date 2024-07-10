package repository

import (
	"database/sql"
	"fmt"
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
	template := fmt.Sprintf(entity.MetadataTableTemplate, r.modelName)

	return createTable(r.db, r.modelName, template)
}

func (r *MetadataRepository) DropTableTemplate() bool {
	return dropTable(r.db, r.modelName)
}

func (r *MetadataRepository) IsTableExist() bool {
	return isTableExist(r.db, r.modelName)
}

func (r *MetadataRepository) Migrate() uint {
	return migrate(r.db, r.modelName, entity.MetadataTableVersion, entity.MetadataTableMigrations)
}
