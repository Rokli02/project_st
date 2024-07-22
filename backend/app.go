package backend

import (
	"context"
	"st/backend/db"
	"st/backend/db/repository"
	"st/backend/model"
	"st/backend/service"
	"st/backend/utils/lang"
	"st/backend/utils/logger"
	"st/backend/utils/settings"
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

	if languageId, has := a.metadatas[settings.MetadataKeys.LanguageId]; !has || languageId.Value == nil {
		logger.ErrorF("Language Id is not loaded into the application!")

		panic(-1)
	} else {
		lang.LoadLanguage(*languageId.Value)
	}
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

	a.SetMetadata(settings.MetadataKeys.CurrentUserId, &model.UpdateMetadata{})
}

func (a *Application) Signup(user *model.SignUpUser) signupResponse {
	err := service.User.SignUp(user)
	if err != nil {
		logger.WarningF("Can't sign up user, (%s)", err)

		return signupResponse{
			Error: &model.ResponseError{
				Code:    400,
				Message: err.Error(),
			},
			Response: lang.Text.User.Get("SIGN_UP_UNSUCCESFUL"),
		}
	}

	return signupResponse{
		Response: lang.Text.User.Get("SIGN_UP_SUCCESFUL"),
	}
}

func (a *Application) GetMetadata(key string) *model.MetadataValue {
	metadata, has := a.metadatas[key]
	if !has {
		return nil
	}

	if metadata.ExpireAt != nil && metadata.ExpireAt.Before(time.Now()) {
		// Remove MetadataValue
		metadata.Value = nil
		metadata.ExpireAt = nil

		service.Metadata.UpdateMetadata(metadata.Id, &model.UpdateMetadata{
			Value:    metadata.Value,
			ExpireAt: metadata.ExpireAt,
		})

		a.metadatas[key] = metadata
	}

	return &metadata
}

func (a *Application) SetMetadata(key string, updateMetadata *model.UpdateMetadata) bool {
	// newMetadata := model.MetadataValue{}
	value, has := a.metadatas[key]
	if !has {
		// HA még nem létezik, akkor CREATE
		createMetadata := &model.MetadataValue{
			Value:    updateMetadata.Value,
			ExpireAt: updateMetadata.ExpireAt,
		}

		if updateMetadata.Type == nil {
			createMetadata.Type = ""
		} else {
			createMetadata.Type = *updateMetadata.Type
		}

		result := service.Metadata.CreateMetadata(key, createMetadata)
		if result == nil {
			return false
		}

		a.metadatas[key] = *result

		return true
	}

	// HA már létezik az elem, akkor UPDATE
	result := service.Metadata.UpdateMetadata(value.Id, updateMetadata)
	if result == nil {
		return false
	}

	a.metadatas[key] = *result

	return true
}

func (a *Application) ReloadLanguage() {
	if languageId, has := a.metadatas[settings.MetadataKeys.LanguageId]; has && languageId.Value != nil {
		lang.LoadLanguage(*languageId.Value)
	}
}

func (a *Application) ReloadLanguageById(languageId string) {
	if a.SetMetadata(settings.MetadataKeys.LanguageId, &model.UpdateMetadata{Value: &languageId}) {
		if languageId, has := a.metadatas[settings.MetadataKeys.LanguageId]; has && languageId.Value != nil {
			lang.LoadLanguage(*languageId.Value)
		}
	}
}

//
//	Types used locally
//

type signupResponse model.StandardResponse[string]
