package controllers

import (
	"context"
	"log"
	"os"
	"yagami/api/models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	Db  *gorm.DB
	CurrentEvent uuid.UUID
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	var err error
	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Error Opening Log: ", err.Error())
	}
	log.SetOutput(file)
	log.Println("\n\n<=================== YAGAMI V0.0.0 STARTED =====================>")
	err = a.SQLStart("./database/yagami.db")
	if err != nil {
		log.Println("Error Starting DB: ", err.Error())
	}
	log.Println("Database started")
}

func Start(a *App) {
}

func (a *App) SQLStart(dbPath string) error {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Println("Error Opening DB:", err.Error())
		return err
	}
	a.Db = db

	err = a.Db.AutoMigrate(&models.User{}, &models.Event{}, &models.Followspot{}, &models.Frame{}, &models.FollowspotCue{}, models.VWXFixture{}, models.Fixture{})
	if err != nil {
		log.Println("Error Migrating DB:", err.Error())
		return err
	}
	return nil
}
