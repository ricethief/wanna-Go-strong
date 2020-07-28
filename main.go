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
var templates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/about.html", "templates/main.html"))

//A Page structure
type Page struct {
	Title string
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

	parsedTemplate, _ := template.ParseFiles("templates/index.html")
	err := parsedTemplate.Execute(w, model.NewVolumeNFrequency(input))
	if err != nil {
		log.Printf("Error occurred while executing the template or writing its output : ", err)
		return
	}
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true) //declare mux router
	myRouter.HandleFunc("/", mainHandler)
	myRouter.HandleFunc("/about", aboutHandler)
	myRouter.HandleFunc("/getVNF/{gender}/{wt}/{ht}/{strength}/{exp}/{age}/{diet}/{sleep}/{stress}/{hw}/{hra}", getvnfHandler)
	myRouter.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	log.Fatal(http.ListenAndServe(":8082", myRouter))

}

func main() {
	handleRequest()
}
