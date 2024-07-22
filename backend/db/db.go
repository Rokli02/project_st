package db

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"st/backend/utils/logger"
	"st/backend/utils/settings"

	_ "modernc.org/sqlite"
)

const CanDeleteDatabase = true

type DB struct {
	Conn         *sql.DB
	dbPath       string
	repositories []Repository
}

func NewDB(path string, models []Repository) *DB {
	db := &DB{
		dbPath:       path,
		repositories: make([]Repository, 0),
	}

	if len(models) > 0 {
		db.repositories = append(db.repositories, models...)
	}

	return db
}

func (db *DB) Connect(ct settings.ConnectTypes) error {
	logger.InfoF("Connecting to database (%s)...", db.dbPath)

	workingDirectory, _ := os.Getwd()
	dbPath := path.Join(workingDirectory, "data", db.dbPath)

	logger.Debug("Database path:", dbPath)

	file, _ := os.OpenFile(dbPath, os.O_RDONLY, os.ModeDevice)
	isDBAlreadyExisted := file != nil
	file.Close()

	// Pre-process before connecting to database
	switch ct {
	case settings.CREATE_IF_NEEDED:
	case settings.CONNECT_IF_EXISTS:
		if !isDBAlreadyExisted {
			return fmt.Errorf("couldn't open database (%s), because it doesn't exist", dbPath)
		}
	case settings.CREATE_ALWAYS:
		if CanDeleteDatabase {
			err := os.Remove(dbPath)
			if err != nil {
				logger.WarningF("Couldn't delete database (%s) for recreation, because (%s)", db.dbPath, err)
			} else {
				isDBAlreadyExisted = false
			}
		}
	}

	// Connection to SQLite Database
	connection, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("couldn't open database (%s)\n\tError:\t%s", dbPath, err.Error())
	}

	db.Conn = connection

	// Enable foreign keys support
	connection.Exec("PRAGMA foreign_keys = ON;")

	logger.Debug("Creating tables that are not in the database already")

	// Attach DB instance to repos AND creating tables if not existing yet
	for _, repo := range db.repositories {
		repo.SetDB(db.Conn)

		if !repo.IsTableExist() {
			logger.InfoF("Creating table '%s'", repo.ModelName())

			repo.CreateTable()
		}
	}

	if !isDBAlreadyExisted {
		logger.Info("Preloading database tables with required records.")
		for _, repo := range db.repositories {
			repo.InitTable()
		}
	}

	return nil
}

func (db *DB) Close() {
	logger.InfoF("Closing Database (%s)...", db.dbPath)

	if db.Conn == nil {
		logger.WarningF("There is no connection to database (%s), the ref is 'nil'", db.dbPath)
		return
	}

	err := db.Conn.Close()
	if err != nil {
		logger.WarningF("Couldn't close database\nError:\t%s", err.Error())
	}
	db.Conn = nil

	for _, repo := range db.repositories {
		repo.SetDB(nil)
	}
}
