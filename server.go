package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type medicine struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	CommercialName string `json:"CommercialName"`
	Concentration float64 `json:"Concentration"`
	Volume int `json:"Volume"`
	Dosage float64 `json:"Dosage"`
}

type groupMedicines []medicine

var medGroup = groupMedicines{
	{
		Id: "1",
		Name: "kefaleksiini",
		CommercialName: "Kefexin",
		Concentration: 100.0,
		Volume: 50,
		Dosage: 50.0,
	},
	{
		Id: "2",
		Name: "kefaleksiini",
		CommercialName: "Kefexin",
		Concentration: 50.0,
		Volume: 100,
		Dosage: 50.0,
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "REST-API server")
}

func getMedicine(w http.ResponseWriter, r *http.Request) {
	medId := mux.Vars(r)["id"]

	for _, med := range medGroup {
		if med.Id == medId {
			json.NewEncoder(w).Encode(med)
		}
	}
}

func getAllMedicines(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(medGroup)
}

func requestHandler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/meds", getAllMedicines).Methods("GET")
	router.HandleFunc("/meds/{id}", getMedicine).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	requestHandler()
}
