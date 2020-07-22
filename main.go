package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"wanna-Go-strong/model"
	"wanna-Go-strong/workoutPkg"

	"github.com/gorilla/mux"
)

//creating Exercise class
func testingit() {
	list1 := list.New()
	list1.PushBack("")
	input := []string{"ff", "Wmiddle", "Hmedium	", "Shigh", "Eadvanced", "A20s", "Dgood", "Sleepgood", "Stresslow", "HW4", "HRA4"}

	workoutPkg.DefaultTraining(input)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page \n")
	modelthing()
	json.NewEncoder(w).Encode(modelthing())
	fmt.Fprintf(w, "\nVisited time", time.Now())
	vars := mux.Vars(r)
	a := vars["a"]
	fmt.Fprintf(w, "\n", a)
	testingit()
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testing Post endpoint")
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
	myRouter.HandleFunc("/{a}", homePage)

	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))

}

func main() {
	handleRequest()
}
