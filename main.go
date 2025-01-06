package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	search "stardust.db/app"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	search := search.NewSearch()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Search Demo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			search,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				UseToolbar: true,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
