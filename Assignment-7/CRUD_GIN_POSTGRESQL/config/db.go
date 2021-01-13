package config

import (
	"log"
	"os"

	"github.com/go-pg/pg"

	controllers "github.com/Ahad-Parmar/Assignment_7/CRUD_GIN_POSTGRESQL/controllers"
)

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "password",
		Addr:     "localhost:5432",
		Database: "golang",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateTodoTable(db)
	controllers.InitiateDB(db)
	return db
}
