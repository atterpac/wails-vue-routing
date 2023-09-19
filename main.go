package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
    // "github.com/wailsapp/wails/v2/pkg/options/linux"
    // "github.com/wailsapp/wails/v2/pkg/options/mac"
    // "github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
    err := wails.Run(&options.App{
        Title:  "alwaysontop",
        Width:  1024,
        Height: 768,
        AssetServer: &assetserver.Options{
          Assets: assets,
        },
        Frameless:       true,
        CSSDragProperty: "widows",
        CSSDragValue:    "1",
		// Color seen while refreshing or resizing
		BackgroundColour:   &options.RGBA{R: 0, G: 0, B: 0, A: 255},
        Bind: []interface{}{
          app,
        },
    })

	if err != nil {
		println("Error:", err.Error())
	}
}
