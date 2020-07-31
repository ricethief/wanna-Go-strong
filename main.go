package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"wanna-Go-strong/model"

	"github.com/gorilla/mux"
)

//compile templates on start
var templates = template.Must(template.ParseGlob("templates/*.html"))

//A Page structure
type Page struct {
	Title string
	Vnf   model.VolumeNFrequency
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

//The handlers.
func mainHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "main", &Page{Title: "Home"})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	display(w, "about", &Page{Title: "About"})
}

func getvnfHandler(w http.ResponseWriter, r *http.Request) {
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
	vnf := model.NewVolumeNFrequency(input)

	display(w, "vnf", &Page{Title: "Volume And Frequency", Vnf: vnf})
}

func vnfformHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		gender := r.FormValue("sexn")
		wt := r.FormValue("weightn")
		ht := r.FormValue("heightn")
		strength := r.FormValue("strengthn")
		exp := r.FormValue("experiencen")
		age := r.FormValue("agen")
		diet := r.FormValue("dietn")
		sleep := r.FormValue("sleepn")
		stress := r.FormValue("stressn")
		hw := r.FormValue("hwn")
		hra := r.FormValue("hran")
		input := []string{gender, wt, ht, strength, exp, age, diet,
			sleep, stress, hw, hra}
		fmt.Println(input)
		vnf := model.NewVolumeNFrequency(input)
		display(w, "vnf", &Page{Title: "Volume And Frequency", Vnf: vnf})
	} else {
		display(w, "vnfform", &Page{Title: "VNF FORM"})
	}
}
func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true) //declare mux router
	myRouter.HandleFunc("/", mainHandler)
	myRouter.HandleFunc("/about", aboutHandler)
	myRouter.HandleFunc("/vnfform", vnfformHandler)
	myRouter.HandleFunc("/getVNF/{gender}/{wt}/{ht}/{strength}/{exp}/{age}/{diet}/{sleep}/{stress}/{hw}/{hra}", getvnfHandler)
	myRouter.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	log.Fatal(http.ListenAndServe(":8082", myRouter))

}

func main() {
	handleRequest()
}
