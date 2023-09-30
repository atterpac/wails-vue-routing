package api

import (
	"context"
	"yagami/api/controllers"
)

func Run() {
	app := controllers.NewApp()	
	app.Startup(context.Background())
}


