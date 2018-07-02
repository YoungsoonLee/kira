package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

// DB ...
// For DB connection
type DB struct {
	Database *mgo.Database
}

// For run once
var _initCtx sync.Once
var _instance *DB

// DBNew ....
// Make DB instance and return
func DBNew() *mgo.Database {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")

	_initCtx.Do(func() {
		_instance = new(DB)

		session, err := mgo.Dial(url)

		if err != nil {
			fmt.Printf("Error en mongo: %+v\n", err)
			os.Exit(1)
		}
		_instance.Database = session.DB("app")
	})
	return _instance.Database
}
