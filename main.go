package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"wanna-Go-strong/model"

	"github.com/gorilla/mux"
)

//creating Exercise class
func testingit() {
	list1 := list.New()
	list1.PushBack("")

}

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page \n")
	testing()
	modelthing()
	json.NewEncoder(w).Encode(modelthing())
	fmt.Fprintf(w, "\nVisited time", time.Now())
	vars := mux.Vars(r)
	a := vars["a"]
	fmt.Fprintf(w, "\n", a)
}

func methpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ahhhhhhhhhhhhhhhhhhhhhhhhhhhhh")
	fmt.Fprint(w, "tt")

}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing Post endpoint")
}
func testing() {
	fmt.Println("testing fun acces")
}

func modelthing() model.Week {
	e := model.NewExercise("Squat", "")
	e2 := model.NewExercise("Bench Press", "")
	e.AddSet(5, 78)
	e.AddSet(5, 78)
	e.AddSet(5, 78)
	e2.AddSet(5, 80)
	e2.AddSet(5, 80)
	e2.AddSet(5, 80)
	d := model.NewDay("Lag day")
	d.AddExercise(e)
	d.AddExercise(e2)
	w := model.NewWeek("HYPERTROPHY")
	w.AddDay("Monday", d)
	return w
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/home/{a}", homePage)

	myRouter.HandleFunc("/add", methpage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))

}

func main() {
	handleRequest()
}
