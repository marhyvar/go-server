package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type Medicine struct {
	Id string `json:"Id"`
	Name string `json:"Name"`
	CommercialName string `json:"CommercialName"`
	Concentration float64 `json:"Concentration"`
	Volume int `json:"Volume"`
	Dosage float64 `json:"Dosage"`
	GroupId string `json:"GroupId"`
}

type GroupMedicines []Medicine

var medGroup = GroupMedicines{
	{
		Id: "1",
		Name: "kefaleksiini",
		CommercialName: "Kefexin",
		Concentration: 100.0,
		Volume: 50,
		Dosage: 50.0,
		GroupId: "1",
	},
	{
		Id: "2",
		Name: "kefaleksiini",
		CommercialName: "Kefexin",
		Concentration: 50.0,
		Volume: 100,
		Dosage: 50.0,
		GroupId: "1",
	},
	{
		Id: "3",
		Name: "amoksisilliini",
		CommercialName: "Amorion",
		Concentration: 100.0,
		Volume: 30,
		Dosage: 40.0,
		GroupId: "2",
	},
	{
		Id: "4",
		Name: "amoksisilliini",
		CommercialName: "Amorion",
		Concentration: 100.0,
		Volume: 50,
		Dosage: 40.0,
		GroupId: "2",
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

func createMedicine(w http.ResponseWriter, r *http.Request) {    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var medicine Medicine 
    json.Unmarshal(reqBody, &medicine)

    medGroup = append(medGroup, medicine)

    json.NewEncoder(w).Encode(medicine)
}

func deleteMedicine(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, medicine := range medGroup {
        if medicine.Id == id {
			//if id matches, removes the medicine from medGroup
            medGroup = append(medGroup[:index], medGroup[index+1:]...)
        }
    }

}

func getByGroupId(w http.ResponseWriter, r *http.Request) {
	groupId := mux.Vars(r)["id"]
	var group []Medicine
	for _, med := range medGroup {
		if med.GroupId == groupId {
			group = append(group, med)
		}
	}
	json.NewEncoder(w).Encode(group)
}

func updateMedicine(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var changedMedicine Medicine

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Something wrong with entered data")
	}
	json.Unmarshal(reqBody, &changedMedicine)

	for index, medicine := range medGroup {
		if medicine.Id == id {
			medicine.Name = changedMedicine.Name
			medicine.CommercialName = changedMedicine.CommercialName
			medicine.Concentration = changedMedicine.Concentration
			medicine.Volume = changedMedicine.Volume
			medicine.Dosage = changedMedicine.Dosage
			medicine.GroupId = changedMedicine.GroupId
			// if id matches, updates medGroup[index] with new value
			medGroup = append(medGroup[:index], medicine)
			json.NewEncoder(w).Encode(medicine)
		}
	}
}

func requestHandler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/medicines", getAllMedicines).Methods("GET")
	router.HandleFunc("/medicines/{id}", getMedicine).Methods("GET")
	router.HandleFunc("/medicines", createMedicine).Methods("POST")
	router.HandleFunc("/medicines/{id}", deleteMedicine).Methods("DELETE")
	router.HandleFunc("/medicines/{id}", updateMedicine).Methods("PUT")
	router.HandleFunc("/groups/{id}/medicines", getByGroupId).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	requestHandler()
}
