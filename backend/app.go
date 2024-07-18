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
	"time"
)

var baseDBRepositories []db.Repository = []db.Repository{repository.User, repository.Metadata}
var userDBRepositories []db.Repository = []db.Repository{}

type Application struct {
	ctx context.Context

	BaseDb    *db.DB
	UserDB    *db.DB
	metadatas model.Metadata
}

func NewApplication() *Application {
	settings.InitSettings()
	repository.InitRepositories()
	service.InitServices()

	return &Application{
		BaseDb:    db.NewDB(settings.Database.BaseDatabaseName, baseDBRepositories),
		metadatas: make(model.Metadata),
	}
}

func (a *Application) Startup(ctx context.Context) {
	a.ctx = ctx
	err := a.BaseDb.Connect(settings.App.BaseDatabaseConnectType)
	if err != nil {
		logger.Error(err.Error())
	}

	a.metadatas = service.Metadata.LoadMetadatas()
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

	userDBPath, err := service.User.Login(user)
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

// TODO:
func (a *Application) Signup(user *model.SignUpUser) {
	message, err := service.User.SignUp(user)
	if err != nil {
		logger.WarningF("Couldn't sign up user, (%s)", err)
	}

	logger.Info(message)
}

func (a *Application) GetMetadata(key string) *model.MetadataValue {
	metadata, has := a.metadatas[key]
	if !has {
		return nil
	}

	if metadata.ExpireAt.Before(time.Now()) {
		// Remove MetadataValue
		metadata.Value = nil
		metadata.ExpireAt = nil

		service.Metadata.UpdateMetadata(metadata.Id, model.UpdateMetadata{
			Value:    metadata.Value,
			ExpireAt: metadata.ExpireAt,
		})

		a.metadatas[key] = metadata
	}

	return &metadata
}

func (a *Application) SetMetadata(key string, value *model.UpdateMetadata) bool {
	// newMetadata := model.MetadataValue{}

	// HA már létezik az elem, akkor UPDATE
	// HA még nem létezik, akkor CREATE
	// metadata, has := a.metadatas[key]
	// if has {
	// 	newMetadata.Id = metadata.Id
	// 	newMetadata.Value = metadata.Value
	// 	newMetadata.Type = metadata.Type
	// 	newMetadata.ExpireAt = metadata.ExpireAt
	// 	// Set new values, to the old metadata
	// 	return false
	// }

	return false
	// Update
}