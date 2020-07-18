package model

import (
	"testing"
)

func TestAddset(t *testing.T) {
	e := NewExercise("Squat", "")
	e.AddSet(5, 5)

	if e.Sets[0] != 5 || e.Sets[1] != 5 {
		t.Error("AddingSet Failed")
	}
}

func TestDeleteset(t *testing.T) {
	e := NewExercise("BenchPress", "")
	e.deleteSet()

	if len(e.Sets) != 0 {
		t.Error("DeleteSet Failed")
	}
}

//testing if exercise is correctly added into day
func TestAddExercise(t *testing.T) {
	e := NewExercise("Squat", "")
	e.AddSet(5, 5)
	d := NewDay("Lag day")
	d.AddDay(e)

	if d.ExerciseList[0].Name != "Squat" {
		t.Error(d)
	}
	//ok i've been strugling with go for first 6 days, but now I kind of like it
}

//testing
