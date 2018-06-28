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

	r.HandleFunc("/echo", routes.Echo).Methods("GET")

	// r.HandleFunc("/images", routes.CreateImages).Methods("POST")
	// r.HandleFunc("/images/{id:[0-9]+}", routes.GetImages).Methods("POST")

	err := http.ListenAndServe(":8080", cors.AllowAll().Handler(r))
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
		log.Fatalln("server start err")
		os.Exit(1)
	}

	log.Println("Listening on port 8080...")

}
