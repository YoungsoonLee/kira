package utils

import (
	"fmt"
	"os"
	"sync"

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
	_initCtx.Do(func() {
		_instance = new(DB)

		session, err := mgo.Dial("mongo:27017") // for docker

		if err != nil {
			fmt.Printf("Error en mongo: %+v\n", err)
			os.Exit(1)
		}
		_instance.Database = session.DB("app")
	})
	return _instance.Database
}
