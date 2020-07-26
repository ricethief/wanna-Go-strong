package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wanna-Go-strong/model"

	"github.com/gorilla/mux"
)

func getVolumNFrequency(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>We Have Volume as jason</h1> \n")
	vars := mux.Vars(r)
	gender := vars["gender"]
	wt := vars["wt"]
	ht := vars["ht"]
	strength := vars["strength"]
	exp := vars["exp"]
	age := vars["age"]
	diet := vars["diet"]
	sleep := vars["sleep"]
	stress := vars["stress"]
	hw := vars["hw"]
	hra := vars["hra"]

	input := []string{gender, wt, ht, strength, exp, age, diet,
		sleep, stress, hw, hra}
	json.NewEncoder(w).Encode(model.NewVolumeNFrequency(input)) //passing json
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing Post endpoint")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/getVNF/{gender}/{wt}/{ht}/{strength}/{exp}/{age}/{diet}/{sleep}/{stress}/{hw}/{hra}", getVolumNFrequency)
	myRouter.HandleFunc("/getvnf/{hra}", getVolumNFrequency)
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))

}

func main() {
	handleRequest()
}
