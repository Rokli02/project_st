package db

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"st/backend/logger"

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

func (db *DB) Connect(ct ConnectTypes) error {
	logger.InfoF("Connecting to database (%s)...", db.dbPath)

	workingDirectory, _ := os.Getwd()
	dbPath := path.Join(workingDirectory, "data", db.dbPath)

	logger.Debug("Database path:", dbPath)

	// Pre-process before connecting to database
	switch ct {
	case CREATE_IF_NEEDED:
		// NOT EXTRA ACTION
	case CREATE_NEW_IF_NOT_EXISTS:
		file, _ := os.OpenFile(dbPath, os.O_RDONLY, os.ModeDevice)
		if file != nil {
			return fmt.Errorf("can't create database (%s)\n\tError: It already exists", dbPath)
		}
	case CONNECT_IF_EXISTS:
		file, _ := os.OpenFile(dbPath, os.O_RDONLY, os.ModeDevice)
		if file == nil {
			return fmt.Errorf("couldn't open database (%s), because it doesn't exist", dbPath)
		}

		file.Close()
	case CREATE_ALWAYS:
		if CanDeleteDatabase {
			os.Remove(dbPath)
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

	logger.Debug("Applying migration changes")

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
