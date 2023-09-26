package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/surrealdb/surrealdb.go"
)

// App struct
type App struct {
	ctx context.Context
	db *surrealdb.DB	
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	fmt.Println("App started")
	a.ctx = ctx
	// go SurrealStart()
	a.db = SurrealConn("ws://127.0.0.1:8069/rpc")
	a.SurrealSignin("root", "root")
	a.Migration()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}


func SurrealStart() {
    ctx := context.Background()
    // Create a new exec command.
    cmd := exec.CommandContext(ctx, "surreal", "start", "--user", "root", "--pass", "root", "file:./database/yagami", "--bind", "0.0.0.0:8069", "&")
    // Start the command.
    err := cmd.Start()
    if err != nil {
        log.Fatalf("Failed to start surreal: %v", err)
    }
	println("SurrealDB started")
}

func (a *App) SurrealSignin(user string, pass string) {
	_, err := a.db.Signin(map[string]interface{}{
		"user": user,
		"pass": pass,
	})
	if err != nil {
		panic(err)
	}

	_, err = a.db.Use("yagami", "yagami")
	if err != nil {
		panic(err)
	}

}

func SurrealConn(db string) *surrealdb.DB{
	conn, err := surrealdb.New(db)
	if err != nil {
		println("Error connecting to SurrealDB")
		panic(err)
	}
	println("Connected to SurrealDB")
	return conn
}


func (a *App) Migration(){
	migration, err := os.ReadFile("./database/migrations/events.surql")
	if err != nil {
		println("Error reading migration")
		panic(err)
	}
	mig_text := string(migration)
	result, err := a.db.Query(mig_text, nil)
	if err != nil {
		println("Error creating table")
		panic(err)
	}
	fmt.Println(result)
}
