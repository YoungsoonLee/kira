package utils

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/YoungsoonLee/kira/api/models"
	"gopkg.in/mgo.v2"
)

// DB ...
// for DB connection
type DB struct {
	Database *mgo.Database
}

// for run once
var _initCtx sync.Once
var _instance *DB

// DBNew ....
// make DB instance and return
func DBNew() *mgo.Database {
	_initCtx.Do(func() {
		_instance = new(DB)
		// session, err := mgo.Dial("mongo:27017") // for docker
		session, err := mgo.Dial("localhost:27017") // for local
		if err != nil {
			fmt.Printf("Error en mongo: %+v\n", err)
			os.Exit(1)
		}
		_instance.Database = session.DB("app")
	})
	return _instance.Database
}

// MakeInitData ...
// make init event data for test
// make 5 test events like below
// Text: event number #1, StartAt: 2018-06-01T00:00:00Z, EndAt: 2018-06-10T23:00:00Z
// Text: event number #2, StartAt: 2018-06-11T00:00:00Z, EndAt: 2018-06-20T23:00:00Z
// Text: event number #3, StartAt: 2018-06-21T00:00:00Z, EndAt: 2018-06-30T23:00:00Z
// Text: event number #4, StartAt: 2018-07-01T00:00:00Z, EndAt: 2018-07-10T23:00:00Z
// Text: event number #5, StartAt: 2018-07-11T00:00:00Z, EndAt: 2018-07-20T23:00:00Z
func MakeInitData() {
	eventDB := DBNew().C("events")
	c, _ := eventDB.Count()

	if c == 0 {
		for i := 0; i < 5; i++ {
			s, _ := time.Parse(time.RFC3339, "2018-06-01T00:00:00Z")
			e, _ := time.Parse(time.RFC3339, "2018-06-10T23:00:00Z")

			if i > 0 {
				s = s.Add(time.Hour * 24 * 10) // add 10 day
				e = e.Add(time.Hour * 24 * 10) // add 10 day
			}

			txt := "event number #" + string(i)

			event := &models.Event{
				Text:    txt,
				StartAt: s,
				EndAt:   e,
			}
			eventDB := DBNew().C("events")
			if err := eventDB.Insert(event); err != nil {
				log.Println(err)
			}
		}
	}
}
