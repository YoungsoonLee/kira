package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/YoungsoonLee/kira/api/models"
	"github.com/YoungsoonLee/kira/api/utils"
	"gopkg.in/mgo.v2/bson"
)

// CreateEvent ...
// Store a new event into DB
// Check. whether there is overlaps event or not, when a new event data insert through the query.
// The meaning of overlaps data is as follows.
// >> If the start date and end date of the event data to be newly input exists
// >> between the start date and end date of the previously input data.
// >> Case 1: start_at of a new event data <= start_at of existing the event data <= end_at of a new event data
// >> Case 2: start_at of a new event data <= end_at of existing the event data <= end_at of a new event data
// >> Case 3: start_at of existing the event data  <= start_at, end_at of a new event data <= end_at of existing the event data
// >. Case 4: new event start_at <= existing event start_at, end_at <= new event end_at
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	// Set time UTC
	time.Local = time.UTC

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
	// Variable for return event data
	result := []models.Event{}

	// Check. whether there is overlaps event or not, when a new event data insert through the query.
	// The meaning of overlaps data is as follows.
	// >> If the start date and end date of the event data to be newly input exists
	// >> between the start date and end date of the previously input data.
	// >> Case 1: start_at of a new event data <= start_at of existing the event data <= end_at of a new event data
	// >> Case 2: start_at of a new event data <= end_at of existing the event data <= end_at of a new event data
	// >> Case 3: start_at of existing the event data  <= start_at, end_at of a new event data <= end_at of existing the event data
	// >. Case 4: new event start_at <= existing event start_at, end_at <= new event end_at
	if err := eventDB.Find(
		bson.M{"$or": []bson.M{
			// Case 1: new event start_at <= existing event start_at <= new evet end_at
			bson.M{"start_at": bson.M{"$gte": event.StartAt, "$lte": event.EndAt}},

			// Case 2: new event start_at <= existing event end_at <= new evet end_at
			bson.M{"end_at": bson.M{"$gte": event.StartAt, "$lte": event.EndAt}},

			// Case 3: existing event start_at <= new event start_at, end_at <= existing event end_at
			bson.M{"$and": []bson.M{
				bson.M{"start_at": bson.M{"$lte": event.StartAt}},
				bson.M{"end_at": bson.M{"$gte": event.EndAt}},
			}},

			// Case 4:  new event start_at <= existing event start_at, end_at <= new event end_at
			bson.M{"$and": []bson.M{
				bson.M{"start_at": bson.M{"$gte": event.StartAt}},
				bson.M{"end_at": bson.M{"$lte": event.EndAt}},
			}},
		}},
	).All(&result); err != nil {
		// Error
		utils.ResponseError(w, err.Error(), nil, http.StatusInternalServerError)
	} else {
		// Check overlaps event data
		if len(result) > 0 {
			// Overlaps
			// Return 400 BedRequestError & overlaps data
			utils.ResponseError(w, "", result, http.StatusBadRequest)
		} else {
			// There is no overlaps data
			// Insert a new event data
			if err := eventDB.Insert(event); err != nil {
				utils.ResponseError(w, err.Error(), nil, http.StatusInternalServerError)
				return
			}

			// Return 200 Success & a new event data
			var re []interface{} // for unmarshall
			re = append(re, event)
			utils.ResponseJSON(w, re)
		}
	}
}

// GetEvents ...
// Return all events data
func GetEvents(w http.ResponseWriter, r *http.Request) {
	// Init model for result
	result := []models.Event{}

	// Connect DB
	eventDB := utils.DBNew().C("events")

	// Query and return all events data
	if err := eventDB.Find(nil).All(&result); err != nil {
		utils.ResponseError(w, err.Error(), nil, http.StatusInternalServerError)
	} else {
		utils.ResponseJSON(w, result)
	}
}
