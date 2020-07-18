package model

import (
	"testing"
)

//testing Exercise
func TestAddset(t *testing.T) {
	e := NewExercise("Squat", "")
	e.AddSet(5, 5)

	if e.Sets[0] != 5 || e.Sets[1] != 5 {
		t.Error("AddingSet Failed")
	}
}

//testing Day
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
	d.AddExercise(e)

	if d.ExerciseList[0].Name != "Squat" {
		t.Error(d)
	}
	//ok i've been strugling with go for first 6 days, but now I kind of like it
}

//testing Week
func TestAddDay(t *testing.T) {
	e := NewExercise("Squat", "")
	e2 := NewExercise("Bench Press", "")
	e.AddSet(5, 78)
	e.AddSet(5, 78)
	e.AddSet(5, 78)
	e2.AddSet(5, 80)
	e2.AddSet(5, 80)
	e2.AddSet(5, 80)
	d := NewDay("Lag day")
	d.AddExercise(e)
	d.AddExercise(e2)
	w := NewWeek("HYPERTROPHY")
	w.AddDay("Monday", d)

	if len(w.Days) != 1 {
		t.Error("AddDay Failed", w.Days["Monday"].DateName)
	}

}
