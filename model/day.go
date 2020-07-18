package model

import (
	"fmt"
)

//Day struct
type Day struct {
	DateName     string
	ExerciseList []Exercise
}

//Constructor
func NewDay(dateName string) Day {
	d := Day{DateName: dateName}
	return d
}

//addExercise
func (d *Day) AddExercise(e Exercise) {
	d.ExerciseList = append(d.ExerciseList, e)

	fmt.Println(e, "added into", d)
}
