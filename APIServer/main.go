package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	dbb "test/api/MongoDB"

	"github.com/gorilla/mux"
)

const URL = "http://localhost:2525/trick-list"

type Trick struct {
	TrickID   string `bson:"_id"`
	TrickName string `bson:"trick_name"`
}

func main() {
	fmt.Println("Starting endpoint on port 2525")

	startServer()
}

func startServer() {
	//Created mux router
	r := mux.NewRouter()

	//end-points I created
	r.HandleFunc("/trick-list", getTricks).Methods("GET")
	r.HandleFunc("/trick-list", addTricks).Methods("POST")
	r.HandleFunc("/trick-list/{trick_id}", deleteTricks).Methods("DELETE")

	// Start server and listen for connections
	http.ListenAndServe(":2525", r)
}

func getTricks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "No-Store")

	trickList, _ := dbb.GetTricks()
	fmt.Fprintf(w, string(trickList))
}

func addTricks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "No-Store")

	response, err := http.Get(URL)
	if err != nil {
		log.Fatal("There was an issue getting: ", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("There was an issue reading the body:", err)
	}

	var t Trick
	json.Unmarshal(data, &t)

	dbb.CreateTricks(t.TrickName)
	fmt.Println("Trick name:", t.TrickName)
}

func deleteTricks(w http.ResponseWriter, r *http.Request) {

}
