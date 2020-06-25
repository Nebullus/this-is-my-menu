package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/gerentes", createGerente).Methods("POST")
	router.HandleFunc("/gerentes", getGerentes).Methods("GET")
	router.HandleFunc("/gerentes/{id}", getGerente).Methods("GET")
	router.HandleFunc("/gerentes/{id}", deleteGerente).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router)))
}

func deleteGerente(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	setCORSRules(w)

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
	setCORSRules(w)

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
	setCORSRules(w)

	gerentesSlice := getGerentesFromDB()
	json.NewEncoder(w).Encode(gerentesSlice)
}

func createGerente(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	setCORSRules(w)
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

func setCORSRules(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}
