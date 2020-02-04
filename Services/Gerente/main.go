package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/gerentes", createGerente).Methods("POST")
	router.HandleFunc("/gerentes", getGerentes).Methods("GET")
	router.HandleFunc("/gerentes/{id}", getGerente).Methods("GET")
	router.HandleFunc("/gerentes/{id}", deleteGerente).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func deleteGerente(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	vars := mux.Vars(r)
	gerente, err := deleteGerenteFromDB(vars["id"])
	if err == nil {
		json.NewEncoder(w).Encode(gerente)
		return
	}
	http.Error(w, err.Error(), 500)
}

func getGerente(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	vars := mux.Vars(r)
	gerente, err := getGerenteFromDB(vars["id"])
	if err == nil {
		json.NewEncoder(w).Encode(gerente)
		return
	}
	http.Error(w, err.Error(), 500)
}

func getGerentes(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	gerentesSlice := getGerentesFromDB()
	json.NewEncoder(w).Encode(gerentesSlice)
}

func createGerente(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var gerente gerente
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newByt := buf.Bytes()
	err := json.Unmarshal(newByt, &gerente)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
		return
	}
	insertGerente(gerente)
}
