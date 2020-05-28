package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "REST-API server")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", router))
}
