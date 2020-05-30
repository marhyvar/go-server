package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
	// Copyright (c) 2012-2018 The Gorilla Authors. All rights reserved.
	// 05/28/2020 gorilla/mux is licensed under the BSD 3-Clause "New" or "Revised" License
)

type Medicine struct {
	Vnr string `json:"Vnr"`
	ActiveIngredient string `json:"ActiveIngredient"`
	CommercialName string `json:"CommercialName"`
	Concentration float64 `json:"Concentration"`
	Volume int `json:"Volume"`
	Atc string `json:"Atc"`
}

type GroupMedicines []Medicine

var medGroup = GroupMedicines{
	{
		Vnr: "398362",
		ActiveIngredient: "kefaleksiini",
		CommercialName: "Kefexin",
		Concentration: 100.0,
		Volume: 50,
		Atc: "J01DB01",
	},
	{
		Vnr: "581900",
		ActiveIngredient: "kefaleksiini",
		CommercialName: "Kefexin",
		Concentration: 50.0,
		Volume: 100,
		Atc: "J01DB01",
	},
	{
		Vnr: "006350",
		ActiveIngredient: "amoksisilliini",
		CommercialName: "Amorion",
		Concentration: 50.0,
		Volume: 100,
		Atc: "J01CA04",
	},
	{
		Vnr: "006299",
		ActiveIngredient: "amoksisilliini",
		CommercialName: "Amorion",
		Concentration: 100.0,
		Volume: 50,
		Atc: "J01CA04",
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "REST-API server")
}

func getMedicine(w http.ResponseWriter, r *http.Request) {
	medVnr := mux.Vars(r)["vnr"]

	for _, med := range medGroup {
		if med.Vnr == medVnr {
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
    Vnr := vars["vnr"]

    for index, medicine := range medGroup {
        if medicine.Vnr == Vnr {
			//if Vnr matches, removes the medicine from medGroup
            medGroup = append(medGroup[:index], medGroup[index+1:]...)
        }
    }

}

func getByAtc(w http.ResponseWriter, r *http.Request) {
	Atc := mux.Vars(r)["atc"]
	var group []Medicine
	for _, med := range medGroup {
		if med.Atc == Atc {
			group = append(group, med)
		}
	}
	json.NewEncoder(w).Encode(group)
}

func updateMedicine(w http.ResponseWriter, r *http.Request) {
	Vnr := mux.Vars(r)["vnr"]
	var changedMedicine Medicine

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Something wrong with entered data")
	}
	json.Unmarshal(reqBody, &changedMedicine)

	for index, medicine := range medGroup {
		if medicine.Vnr == Vnr {
			medicine.ActiveIngredient = changedMedicine.ActiveIngredient
			medicine.CommercialName = changedMedicine.CommercialName
			medicine.Concentration = changedMedicine.Concentration
			medicine.Volume = changedMedicine.Volume
			medicine.Atc = changedMedicine.Atc
			// if Vnr matches, updates medGroup[index] with new value
			medGroup = append(medGroup[:index], medicine)
			json.NewEncoder(w).Encode(medicine)
		}
	}
}

func requestHandler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/medicines", getAllMedicines).Methods("GET")
	router.HandleFunc("/medicines/{vnr}", getMedicine).Methods("GET")
	router.HandleFunc("/medicines", createMedicine).Methods("POST")
	router.HandleFunc("/medicines/{vnr}", deleteMedicine).Methods("DELETE")
	router.HandleFunc("/medicines/{vnr}", updateMedicine).Methods("PUT")
	router.HandleFunc("/groups/{atc}/medicines", getByAtc).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	requestHandler()
}
