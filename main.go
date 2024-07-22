package main

import (
	"context"
	"embed"
	"st/backend"
	"st/backend/model"
	"st/backend/utils"
	myLogger "st/backend/utils/logger"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func _() {
	// Create an instance of the app structure
	app := backend.NewApplication()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Seridet",
		Width:            1024,
		Height:           768,
		AssetServer:      &assetserver.Options{Assets: assets},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		OnShutdown:       app.Shutdown,
		Windows: &windows.Options{
			DisablePinchZoom: true,
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func main() {
	myLogger.Info("START")
	app := backend.NewApplication()
	app.Startup(context.Background())

	// sRes := repository.User.Save(&entity.User{Login: "Bruh", Password: "abc", Name: utils.ToRef("Bakkancs Brádör")})
	// myLogger.Debug("Saved Result:", sRes)

	// fRes1 := repository.User.FindOneByLoginAndPassword("Bruh", "abc")
	// myLogger.Debug("Find One Result:", fRes1)
	// if fRes1.Name != nil {
	// 	myLogger.Debug("Name:", *fRes1.Name)
	// }

	// myLogger.Debug(repository.Metadata.FindAll())
	myLogger.Error("___Választó vonal___")

	myLogger.Debug("Sign up response:", app.Signup(&model.SignUpUser{Login: "brotha", Password: "123", Name: utils.ToRef("Vér Testvér")}))
	myLogger.Debug("LanguageId Metadata:", app.GetMetadata("languageId"))
	app.ReloadLanguageById("hu")
	myLogger.Debug("LanguageId Metadata:", app.GetMetadata("languageId"))
	myLogger.Debug("Sign up response:", app.Signup(&model.SignUpUser{Login: "brotha", Password: "123", Name: utils.ToRef("Vér Testvér")}))

	defer app.BaseDb.Close()
}
