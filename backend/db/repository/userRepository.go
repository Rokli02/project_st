package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"st/backend/db"
	"st/backend/db/entity"
	"st/backend/logger"
	"st/backend/settings"
	"time"
)

type UserRepository struct {
	db        *sql.DB
	modelName string
}

var _ db.Repository = (*UserRepository)(nil)

func (r *UserRepository) FindOneByLoginAndPassword(loginParam, passwd string) *entity.User {
	template := fmt.Sprintf("SELECT id, login, name, db_path FROM %s WHERE login = ? AND password = ?;", r.modelName)
	row := r.db.QueryRow(template, loginParam, passwd)

	u := &entity.User{}

	if err := row.Scan(&(u.Id), &(u.Login), &(u.Name), &(u.DBPath)); err != nil {
		logger.WarningF("Can't return with %s (%v)", r.modelName, err)

		return nil
	}

	return u
}

func (r *UserRepository) IsExist(login string) bool {
	qTemplate := fmt.Sprintf("SELECT count(*) FROM %s WHERE login = ?;", r.modelName)

	row := r.db.QueryRow(qTemplate, login)
	count := 0

	if err := row.Scan(&count); err != nil || count == 0 {
		logger.DebugF("User is already in table %s, Error(%v)", r.modelName, err)

		return false
	}

	return true
}

func (r *UserRepository) Save(user *entity.User) bool {
	// Generate a DB file path for user inside 'data/usr/'
	user.DBPath = createUsrDBPath(user.Login)
	if !isUsrDBPathAvailable(user.DBPath) {
		logger.WarningF("Database with such name is already in use('%s'). Try again a bit later!", user.DBPath)

		return false
	}

	logger.InfoF("Created a new database with name '%s'", user.DBPath)

	mTemplate := fmt.Sprintf("INSERT INTO %s (login, password, name, db_path, created_at) VALUES (?,?,?,?,?);", r.modelName)

	res, err := r.db.Exec(mTemplate, user.Login, user.Password, user.Name, user.DBPath, time.Now().Format(settings.Database.DateFormat))
	if err != nil {
		logger.ErrorF("Can't save to %s table, because: (%v)", r.modelName, err)

		return false
	}

	inserted, _ := res.LastInsertId()
	if inserted == 0 {
		logger.ErrorF("Can't save to %s table, because: (%v)", r.modelName, err)

		return false
	}

	return true
}

func createUsrDBPath(loginParam string) string {
	login := []rune(loginParam)
	currentTimeStamp := fmt.Sprintf("%d", time.Now().UnixMilli())

	loginFirstLetter := string(login[0])
	loginMiddleLetter := string(login[len(login)>>1])
	loginLastLetter := string(login[len(login)-1])

	lhCurrentTimeStamp := currentTimeStamp[0 : len(currentTimeStamp)>>2]
	mhCurrentTimeStamp := currentTimeStamp[len(currentTimeStamp)>>2 : len(currentTimeStamp)>>1+len(currentTimeStamp)>>2]
	fhCurrentTimeStamp := currentTimeStamp[len(currentTimeStamp)>>1+len(currentTimeStamp)>>2:]

	return path.Join("usr", fmt.Sprintf(
		"%s%s-%s%s-%s%s",
		loginFirstLetter, fhCurrentTimeStamp, loginMiddleLetter, mhCurrentTimeStamp, loginLastLetter, lhCurrentTimeStamp),
	)
}

func isUsrDBPathAvailable(personalDBPath string) bool {
	workDir, _ := os.Getwd()
	file, _ := os.OpenFile(path.Join(workDir, "data", "user", personalDBPath), os.O_RDONLY, os.ModeDevice)
	return file == nil
}

//
//	db.Repository Implementations
//

func (r *UserRepository) SetDB(db *sql.DB) {
	r.db = db
}

func (r *UserRepository) ModelName() string {
	return r.modelName
}

func (r *UserRepository) CreateTable() bool {
	return createTable(r.db, &entity.User{}, r.modelName)
}

func (r *UserRepository) DropTable() bool {
	return dropTable(r.db, r.modelName)
}

func (r *UserRepository) IsTableExist() bool {
	return isTableExist(r.db, r.modelName)
}

func (r *UserRepository) Migrate() uint {
	return migrate(r.db, r.modelName, &entity.User{})
}

func (r *UserRepository) InitTable() {}
