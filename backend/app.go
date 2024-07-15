package backend

import (
	"context"
	"st/backend/db"
	"st/backend/db/repository"
	"st/backend/logger"
	"st/backend/model"
	"st/backend/service"
	"st/backend/settings"
	"strings"
)

var baseDBRepositories []db.Repository = []db.Repository{repository.UserRepo, repository.MetadataRepo}
var userDBRepositories []db.Repository = []db.Repository{}

type Application struct {
	ctx context.Context

	BaseDb    *db.DB
	UserDB    *db.DB
	Metadatas map[string]string
}

func NewApplication() *Application {
	settings.InitSettings()
	repository.InitRepositories()
	service.InitServices()

	return &Application{
		BaseDb:    db.NewDB(settings.Database.BaseDatabaseName, baseDBRepositories),
		Metadatas: make(map[string]string),
	}
}

func (a *Application) Startup(ctx context.Context) {
	a.ctx = ctx
	err := a.BaseDb.Connect(db.CREATE_ALWAYS)
	if err != nil {
		logger.Error(err.Error())
	}

	// TODO: Get Metadatas From DB
}

func (a *Application) Shutdown(ctx context.Context) {
	a.ctx = ctx
	a.BaseDb.Close()

	if a.UserDB != nil {
		a.UserDB.Close()
	}
}

// Nem tudom mik legyenek a paraméterek
func (a *Application) Login(user *model.LoginUser) {
	if user == nil || strings.TrimSpace(user.Login) == "" || strings.TrimSpace(user.Password) == "" {
		return
	}

	userDBPath, err := service.UserServ.Login(user)
	if err != nil {
		logger.WarningF("something happened during logging in, %s", err.Error())

		return
	}

	// TODO: Update Metadata keys

	a.UserDB = db.NewDB(userDBPath, userDBRepositories)
}

func (a *Application) Logout() {
	a.UserDB = nil

	// TODO: Update Metadata keys
}
