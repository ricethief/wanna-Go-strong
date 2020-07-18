package model

import (
	"fmt"
)

//Exercise struct
type Exercise struct {
	Name      string    //name of the exercise
	VideoLink string    //url for videolink
	Sets      []float64 //set of Exercise[reps][intensity]
}

//constructor
func NewExercise(name, videoLink string) Exercise {

	var SetsArray = []float64{}
	//this constructor assign SetArray as a Sets[][]float64
	e := Exercise{Name: name, VideoLink: videoLink, Sets: SetsArray}
	return e
}

//addSet adds new rap,intensity into the Exercise
func (e *Exercise) AddSet(_rep float64, _intensity float64) {
	rep := _rep             //rep array
	intensity := _intensity //intensity array
	e.Sets = append(e.Sets, rep)
	e.Sets = append(e.Sets, intensity)

	fmt.Println("set added", e.Sets)
}

//deleteing a set from the behind
//return deleted array do nothing if array is empty
func (e *Exercise) deleteSet() {
	if len(e.Sets) == 0 {
		fmt.Println("Nothing to Delete here :)")
	} else {
		e.Sets = e.Sets[:len(e.Sets)-2]
	}
}
