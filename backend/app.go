package backend

import (
	"context"
	"fmt"
	"st/backend/db"
	"st/backend/db/repository"
	"st/backend/model"
	"st/backend/service"
	"st/backend/utils"
	"st/backend/utils/lang"
	"st/backend/utils/logger"
	"st/backend/utils/settings"
	"strings"
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
	// Connect to base database
	a.ctx = ctx
	err := a.BaseDb.Connect(settings.App.BaseDatabaseConnectType)
	if err != nil {
		logger.Error(err.Error())
	}

	// Load metadatas and texts
	a.metadatas = service.Metadata.LoadMetadatas()

	if languageId, has := a.metadatas[settings.MetadataKeys.LanguageId]; !has || languageId.Value == nil {
		logger.ErrorF("Language Id is not loaded into the application!")

		panic(-1)
	} else {
		lang.LoadLanguage(*languageId.Value)
	}

	// Check if there is an user that stayed logged in AND connects to its DB
	if a.HasLoggedInUser(nil) {
		a.connectUserDB()
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
func (a *Application) Login(loginUser *model.LoginUser) bool {
	if loginUser == nil || strings.TrimSpace(loginUser.Login) == "" || strings.TrimSpace(loginUser.Password) == "" {
		return false
	}

	user, err := service.User.Login(loginUser)
	if err != nil || user == nil {
		logger.WarningF("something happened during logging in, %s", err)

		return false
	}

	// Update currentUser in Metadata
	a.SetMetadata(settings.MetadataKeys.CurrentUserId, &model.UpdateMetadata{Value: utils.ToRef(fmt.Sprintf("%d", user.Id))})

	if a.HasLoggedInUser(user) {
		return a.connectUserDB()
	}

	return false
}

func (a *Application) Logout() {
	if a.UserDB != nil {
		a.UserDB.Close()
	}

	a.UserDB = nil

	a.SetMetadata(settings.MetadataKeys.CurrentUserId, &model.UpdateMetadata{Value: utils.ToRef("")})
}

func (a *Application) Signup(user *model.SignUpUser) model.StandardResponse {
	err := service.User.SignUp(user)
	if err != nil {
		logger.WarningF("Can't sign up user, (%s)", err)

		return model.StandardResponse{
			Error: &model.ResponseError{
				Code:    400,
				Message: err.Error(),
			},
			Response: lang.Text.User.Get("SIGN_UP_UNSUCCESFUL"),
		}
	}

	return model.StandardResponse{
		Response: lang.Text.User.Get("SIGN_UP_SUCCESFUL"),
	}
}

func (a *Application) GetMetadata(key string) *model.MetadataValue {
	metadata, has := a.metadatas[key]
	if !has {
		return nil
	}

	if utils.IsDateExpired(metadata.ExpireAt) {
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

func (a *Application) HasLoggedInUser(user *model.User) bool {
	if a.UserDB != nil {
		return true
	}

	var currentUser *model.User = user

	if currentUser == nil {
		// Check if there is an user that stayed logged in
		currentUserMetadata := a.GetMetadata(settings.MetadataKeys.CurrentUserId)
		if currentUserMetadata == nil || currentUserMetadata.Value == nil {
			return false
		}

		// Parsing to an unsigned integer is totally fine, because ID must be bigger than 0 AND ParseInt would call this too
		var currentUserId int64
		if read, err := fmt.Sscan(*currentUserMetadata.Value, &currentUserId); read != 1 || err != nil {
			logger.WarningF("Error occured during parsing current user id (%s)", err)

			return false
		}

		currentUser = service.User.FindById(currentUserId)
		if currentUser == nil {
			return false
		}
	}

	// If a user is found, then connect to DB
	a.UserDB = db.NewDB(*currentUser.DBPath, userDBRepositories)

	return true
}

func (a *Application) connectUserDB() bool {
	if err := a.UserDB.Connect(settings.CREATE_ALWAYS); err != nil {
		logger.ErrorF("Couldn't open user's database (%s). Raised error: %s", a.UserDB.DBPath, err)

		return false
	}

	return true
}
