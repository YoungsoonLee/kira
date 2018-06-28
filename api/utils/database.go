package utils

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/YoungsoonLee/mrecun/api/models"
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
		session, err := mgo.Dial("mongo:27017")
		if err != nil {
			fmt.Printf("Error en mongo: %+v\n", err)
			os.Exit(1)
		}
		_instance.Database = session.DB("app")
	})
	return _instance.Database
}

// MakeInitData ...
// make init data
func MakeInitData() {
	imageDB := DBNew().C("images")
	c, _ := imageDB.Count()

	if c == 0 {
		// make init data
		// promotion_id =1 && start_at: 2018-01-01T00:00:00Z, end_at: 2019-12-31T23:00:00Z,  data 5개
		for i := 0; i < 5; i++ {
			s, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00Z")
			e, _ := time.Parse(time.RFC3339, "2019-12-31T23:00:00Z")

			image := &models.Image{
				URL:         "https://static-mercariapp-com.akamaized.net/thumb/photos/m18459967832_1.jpg",
				PromotionID: 1,
				StartAt:     s,
				EndAt:       e,
			}
			imageDB := DBNew().C("images")
			if err := imageDB.Insert(image); err != nil {
				log.Println(err)
			}
		}

		// promotion_id =1 && start_at: 2017-01-01T00:00:00Z, end_at: 2018-06-01T23:00:00Z,  data 5개
		for i := 0; i < 5; i++ {
			s, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
			e, _ := time.Parse(time.RFC3339, "2018-06-01T23:00:00Z")

			image := &models.Image{
				URL:         "https://static-mercariapp-com.akamaized.net/thumb/photos/m18459967832_1.jpg",
				PromotionID: 1,
				StartAt:     s,
				EndAt:       e,
			}
			imageDB := DBNew().C("images")
			if err := imageDB.Insert(image); err != nil {
				log.Println(err)
			}
		}

		// promotion_id =2 && start_at: 2018-01-01T00:00:00Z, end_at: 2019-12-31T23:00:00Z,  data 5개
		for i := 0; i < 3; i++ {
			s, _ := time.Parse(time.RFC3339, "2018-01-01T00:00:00Z")
			e, _ := time.Parse(time.RFC3339, "2019-12-31T23:00:00Z")

			image := &models.Image{
				URL:         "https://static-mercariapp-com.akamaized.net/thumb/photos/m18459967832_1.jpg",
				PromotionID: 2,
				StartAt:     s,
				EndAt:       e,
			}
			imageDB := DBNew().C("images")
			if err := imageDB.Insert(image); err != nil {
				log.Println(err)
			}
		}

		// promotion_id =2 && start_at: 2017-01-01T00:00:00Z, end_at: 2018-06-01T23:00:00Z,  data 5개
		for i := 0; i < 3; i++ {
			s, _ := time.Parse(time.RFC3339, "2017-01-01T00:00:00Z")
			e, _ := time.Parse(time.RFC3339, "2018-06-01T23:00:00Z")

			image := &models.Image{
				URL:         "https://static-mercariapp-com.akamaized.net/thumb/photos/m18459967832_1.jpg",
				PromotionID: 2,
				StartAt:     s,
				EndAt:       e,
			}
			imageDB := DBNew().C("images")
			if err := imageDB.Insert(image); err != nil {
				log.Println(err)
			}
		}
	}

}

// MakeInitIP ...
// make init IP
func MakeInitIP() {
	ipDB := DBNew().C("ips")
	c, _ := ipDB.Count()

	if c == 0 {
		ip := &models.IP{"10.0.0.1", true}
		if err := ipDB.Insert(ip); err != nil {
			log.Println(err)
		}

		ip = &models.IP{"10.0.0.2", true}
		if err := ipDB.Insert(ip); err != nil {
			log.Println(err)
		}
	}
}
