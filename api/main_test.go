package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/YoungsoonLee/kira/api/models"
	"github.com/YoungsoonLee/kira/api/routes"
	"github.com/YoungsoonLee/kira/api/utils"
	"github.com/gorilla/mux"
)

// TestAddEvent ...
// Check overlaps event data with qiery when a new event data insert
// That is an overlaps data.
// >> If the start date and end date of the event data to be newly input exists
// >> between the start date and end date of the previously input data.
// >> Case 1: new event start_at <= existing event start_at <= new evet end_at
// >> Case 2: new event start_at <= existing event end_at <= new evet end_at
// >> Case 3: existing event start_at <= new event start_at, end_at <= existing event end_at
// >> Case 4:  new event start_at <= existing event start_at, end_at <= new event end_at
// Test Case 1: Test add a event when there is overlpas events
// >> Test Date: start_at: 2018-06-04, end_at: 2018-06-22
// >> Expect: return 400 Bad Request & three overlaps data
func TestOverlapsEvent(t *testing.T) {
	// make init db
	err := makeInitData()
	if err != nil {
		t.Error(err)
	}

	// Test case #1.
	// Test add a event when there is overlpas events
	// Expect: return 400 Bad Request & three overlaps data
	t.Run("Test add a event when there is overlpas events", func(t *testing.T) {
		// Set init test date: start_at: 2018-06-04, end_at: 2018-06-22
		s := time.Date(2018, 06, 04, 00, 00, 00, 00, time.UTC)
		e := time.Date(2018, 06, 22, 00, 00, 00, 00, time.UTC)

		// Make payload data JSON
		payload := models.Event{Text: "a new kira event", StartAt: s, EndAt: e}
		p, err := json.Marshal(payload)

		// Make test request
		req, err := http.NewRequest("POST", "/event", bytes.NewBuffer(p))
		if err != nil {
			t.Error(err)
		}
		// Set Header for JSON
		req.Header.Set("Content-Type", "application/json")

		// Send test data
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/event", routes.CreateEvent)
		router.ServeHTTP(rr, req)

		// Get response
		responseData, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Errorf("read body error")
		}

		// Unmarshall responseData
		var decoded []interface{}
		if err := json.Unmarshal(responseData, &decoded); err != nil {
			t.Error("Unmarshal error: ", err)
		}

		expect := 3 // expexte overlaps data count
		got := len(decoded)

		// Check result
		if got != expect && rr.Code != 400 {
			t.Errorf("got '%d', expect '%d'", got, expect)
		}
	})
}

// Test Case 2: start_at: 2018-09-01, end_at: 2018-09-10
// >> Expect: return 200 Success & a new event data
// If you want to test the same below test data for overlap,
// you use TestOverlapsEvent. It's just for testing add a new event data
func TestAddNewEvent(t *testing.T) {
	// Test case #2.
	// Test add a new event data when there is no overlpas events
	// Test data: start_at: 2018-10-01, end_at: 2018-10-10
	// >> Expect: return 200 Success & a new event data
	t.Run("Test add a new event when there is no overlpas events", func(t *testing.T) {
		// Set init test date: start_at: 2018-06-04, end_at: 2018-06-22
		s := time.Date(2018, 10, 01, 00, 00, 00, 00, time.UTC)
		e := time.Date(2018, 10, 10, 00, 00, 00, 00, time.UTC)

		// Make payload data JSON
		payload := models.Event{Text: "a new kira event in Oct", StartAt: s, EndAt: e}
		p, err := json.Marshal(payload)

		// Make test request
		req, err := http.NewRequest("POST", "/event", bytes.NewBuffer(p))
		if err != nil {
			t.Error(err)
		}
		// Set Header for JSON
		req.Header.Set("Content-Type", "application/json")

		// Send test data
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/event", routes.CreateEvent)
		router.ServeHTTP(rr, req)

		// Get response
		responseData, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Errorf("read body error")
		}

		// Unmarshall responseData
		var decoded []interface{}
		if err := json.Unmarshal(responseData, &decoded); err != nil {
			t.Error("Unmarshal error: ", err)
		}

		expect := 1 // expexte a new data count
		got := len(decoded)

		// Check result
		if got != expect && rr.Code != 200 {
			t.Errorf("got '%d', expect '%d', Code '%d'", got, expect, rr.Code)
		}
	})
}

// MakeInitData ...
// make init event data for test
// make 5 test events like below
// Text: event number #1, StartAt: 2018-06-01T00:00:00Z, EndAt: 2018-06-10T23:00:00Z
// Text: event number #2, StartAt: 2018-06-11T00:00:00Z, EndAt: 2018-06-20T23:00:00Z
// Text: event number #3, StartAt: 2018-06-21T00:00:00Z, EndAt: 2018-06-30T23:00:00Z
// Text: event number #4, StartAt: 2018-07-01T00:00:00Z, EndAt: 2018-07-10T23:00:00Z
// Text: event number #5, StartAt: 2018-07-11T00:00:00Z, EndAt: 2018-07-20T23:00:00Z
func makeInitData() error {
	// Connect DB
	eventDB := utils.DBNew().C("events")

	// Remove all documents
	// eventDB.RemoveAll(nil)

	// Check count for alredy having data
	c, _ := eventDB.Count()

	// Init start day
	sd := 01
	// Init end day
	ed := 10

	if c == 0 {
		for i := 0; i < 5; i++ {
			// Add 10days
			if i > 0 {
				sd = sd + 10
				ed = ed + 10
			}

			// Set init time
			s := time.Date(2018, 06, sd, 00, 00, 00, 00, time.UTC)
			e := time.Date(2018, 06, ed, 00, 00, 00, 00, time.UTC)

			// Make text
			txt := "event number #" + strconv.Itoa(i)

			// Make data
			event := &models.Event{
				Text:    txt,
				StartAt: s,
				EndAt:   e,
			}

			// Insert DB
			eventDB := utils.DBNew().C("events")
			if err := eventDB.Insert(event); err != nil {
				log.Println(err)
				return err
			}
		}

	}

	return nil
}
