package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start start a http API
func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/api/time", TimeInAnotherTimeZone).Methods(http.MethodGet)

	fmt.Println("Listening API on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
