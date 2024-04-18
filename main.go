package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"teletraffic/backend/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go func() {
		e := echo.New()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))

		routes.TrafficRoutesInit(e)
		routes.NetworkRoutesInit(e)
		routes.RFRoutesInit(e)
		if err := e.Start(":5174"); err != nil {
			println("Echo server error:", err.Error())
		}
	}()
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options

	err := wails.Run(&options.App{
		Title:            "TTG International - Teletraffic Mate (New)",
		DisableResize:    false,
		WindowStartState: options.Maximised,

		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}

}
