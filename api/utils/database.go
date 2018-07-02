package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/tkanos/gonfig"
	"gopkg.in/mgo.v2"
)

// Configuration ...
// For config
type Configuration struct {
	DbURL string
}

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

	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		log.Fatalln(err)
	}

	_initCtx.Do(func() {
		_instance = new(DB)

		session, err := mgo.Dial(configuration.DbURL)

		if err != nil {
			fmt.Printf("Error en mongo: %+v\n", err)
			os.Exit(1)
		}
		_instance.Database = session.DB("app")
	})
	return _instance.Database
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), "config", strings.Join(filename, ""))

	return filePath
}
