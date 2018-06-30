package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/YoungsoonLee/kira/api/models"
	"github.com/YoungsoonLee/kira/api/utils"
	"gopkg.in/mgo.v2/bson"
)

// CreateEvent ...
// store event information into DB
// Check overlaps event data with qiery when new event insert
// That is an overlaps data.
// If the start date and end date of the event data to be newly input exists
// between the start date and end date of the previously input data.
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Read body from request
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ResponseError(w, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// Init & Read event data
	event := &models.Event{}
	err = json.Unmarshal(data, event)
	if err != nil {
		utils.ResponseError(w, err.Error(), nil, http.StatusBadRequest)
		return
	}

	// Init DB
	eventDB := utils.DBNew().C("events")
	// variable for return event data
	result := []models.Event{}

	// Check overlaps event data with qiery when new event insert
	// That is an overlaps data.
	// If the start date and end date of the event data to be newly input exists
	// between the start date and end date of the previously input data.
	if err := eventDB.Find(
		bson.M{
			"start_at": bson.M{"$gte": event.StartAt},
			"end_at":   bson.M{"$lte": event.EndAt},
		},
	).All(&result); err != nil {
		// Error
		utils.ResponseError(w, err.Error(), nil, http.StatusInternalServerError)
	} else {
		// Check overlaps event data
		if len(result) > 0 {
			// overlaps
			utils.ResponseError(w, "event overlaps", result, http.StatusBadRequest)
		} else {
			// There is no overlaps data
			// Insert new event info
			if err := eventDB.Insert(event); err != nil {
				utils.ResponseError(w, err.Error(), nil, http.StatusInternalServerError)
				return
			}

			// Return new event information
			utils.ResponseJSON(w, "created event", event)
		}
	}
}

// GetEvents ...
// return all events data
func GetEvents(w http.ResponseWriter, r *http.Request) {
	// init model for result
	result := []models.Event{}

	// connect DB
	eventDB := utils.DBNew().C("events")

	// query and return all events data
	if err := eventDB.Find(nil).All(&result); err != nil {
		utils.ResponseError(w, err.Error(), nil, http.StatusInternalServerError)
	} else {
		utils.ResponseJSON(w, nil, result)
	}
}
