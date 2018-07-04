package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/YoungsoonLee/kira/api/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Set up routes
	r := mux.NewRouter()

	// route echo for healty check
	r.HandleFunc("/echo", routes.Echo).Methods("GET")

	// route add event
	// if there is overlaps data, return message and overlaps data when a new event add
	r.HandleFunc("/event", routes.CreateEvent).Methods("POST")

	// route get all events data
	r.HandleFunc("/event", routes.GetEvents).Methods("GET")

	log.Println("Listening on port 8080...")

	// start server
	err := http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
		log.Fatalln("server start err")
		os.Exit(1)
	}

}
