package repository

import (
	"database/sql"
	"fmt"
	"st/backend/db"
	"st/backend/logger"
	"strings"
)

func createTable(db *sql.DB, modelName string, template string) bool {
	res, err := db.Exec(template)
	if err != nil {
		logger.ErrorF("Error occured during creating table '%s' (%v)", modelName, err)

		return false
	}

	affected, err := res.RowsAffected()
	if err != nil {
		logger.WarningF("For some reason there are no affected rows for creating '%s' table", modelName)
	}

	return affected > 0
}

func dropTable(db *sql.DB, modelName string) bool {
	template := fmt.Sprintf("DROP TABLE %s;", modelName)

	res, _ := db.Exec(template)
	affacted, err := res.RowsAffected()

	if err != nil || affacted < 1 {
		logger.WarningF("Can't drop table (%v)", err)

		return false
	}

	return true
}

func isTableExist(db *sql.DB, modelName string) bool {
	template := fmt.Sprintf("SELECT count(*) from sqlite_master WHERE type = 'table' AND name = '%s';", modelName)
	var count int64 = -1

	if res, _ := db.Query(template); res.Next() {
		res.Scan(&count)
		res.Close()
	}

	if count < 1 {
		return false
	}

	return true
}

func migrate(db *sql.DB, modelName string, versionUpTo uint, migrations []db.Migration) uint {
	sb := strings.Builder{}

	for _, mig := range migrations {
		if mig.Version >= versionUpTo {
			sb.WriteString(mig.Template)
			sb.WriteString("\n")
		}
	}

	res, err := db.Exec(sb.String())
	if err != nil {
		logger.ErrorF("Somer error occured during migrating table '%s' to a newer version (%v)", modelName, err)

		return 0
	}

	affacted, err := res.RowsAffected()
	if err != nil {
		logger.WarningF("Error occured in '%s' migration", modelName)
	}

	return uint(affacted)
}
