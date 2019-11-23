package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// main function
func main() {

	feedDBwVideo()
	router := mux.NewRouter()
	router.HandleFunc("/gerentes", GetG).Methods("GET")
	/**
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	*/
	log.Fatal(http.ListenAndServe(":8000", router))

}

// "GetVideo to get all videos from DB"
func GetVideo(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	v := getVideoFromDB()
	json.NewEncoder(w).Encode(v)

}

func PostGerente(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
