package main

import (
	"embed"
	"yagami/api"
	"yagami/api/controllers"

	"github.com/wailsapp/wails/v2"

	// "github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	// "github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := controllers.NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Project Yagami",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:       true,
		CSSDragProperty: "widows",
		CSSDragValue:    "1",
		OnStartup: app.Startup,
		Logger: nil,
		// Color seen while refreshing or resizing
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 255},
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			BackdropType:                      windows.Mica,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
		},
		Linux: &linux.Options{
            WindowIsTranslucent: false,
            WebviewGpuPolicy: linux.WebviewGpuPolicyAlways,
            ProgramName: "wails",
        },
	},
	)

	if err != nil {
		println("Error:", err.Error())

	}
	api.Run()	
}
