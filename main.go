package main

import (
	"context"
	"embed"
	"st/backend"
	"st/backend/db/entity"
	"st/backend/db/repository"
	"st/backend/logger"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
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
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func main() {
	app := backend.NewApplication()
	app.Startup(context.Background())

	sRes := repository.UserRepo.Save(&entity.User{Login: "Bruh", Password: "abc"})
	logger.Info("Saved Result:", sRes)

	fRes1 := repository.UserRepo.FindOneByLoginAndPassword("Bruh", "abc")
	logger.Info("Find One Result:", fRes1)

	defer app.BaseDb.Close()
}
