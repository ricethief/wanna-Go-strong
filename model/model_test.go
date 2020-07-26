package model

import (
	"fmt"
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

//Testing GetMev
func TestGetMEV(t *testing.T) {
	//	input := []string{"ff", "Wmiddle", "Hmedium", "Smoderate", "Eintermediate",
	//		"A20s", "Davg", "Sleepavg", "Stresslow", "HW2", "HRA2"}
	//		input2 := []string{"mm", "Wmiddle", "Hmedium", "Smoderate", "Eintermediate",
	//		"A30s", "Dgood", "Sleepgood", "Stresslavg", "HW3", "HRA3"}

	input3 := []string{"mm", "Wmiddle", "Hmedium", "Shigh", "Eadvanced",
		"A20s", "Dgood", "Sleepavg", "Stresslavg", "HW4", "HRA3"}
	DefaultTraining(input3)
}

func TestMarisaExampleMRV(t *testing.T) {
	input := []string{"ff", "Wmiddle", "Hmedium", "Smoderate", "Eintermediate",
		"A20s", "Davg", "Sleepavg", "Stresslow", "HW2", "HRA2"}
	s, b, d := GetMRV(input)

	if s[0] != 21 {
		fmt.Println(s[0], "is squat mrv, rest :", b, d)
		t.Error("Squat Mrv incorrect expected 21 but it was", s[0])
	}
	if b[0] != 24 {
		fmt.Println(b[0], "is squat mrv, rest :", b, d)
		t.Error("Benchpress Mrv incorrect expected 24 but it was", b[0])
	}
	if d[0] != 18 {
		fmt.Println(d[0], "is squat mrv, rest :", b, d)
		t.Error("Deadlift Mrv incorrect expected 18 but it was", d[0])
	}

}

func TestGetFrequency(t *testing.T) {
	input := []string{"ff", "Wmiddle", "Hmedium", "Smoderate", "Eintermediate",
		"A20s", "Davg", "Sleepavg", "Stresslow", "HW2", "HRA2"}
	s, b, d := GetFreqiency(input)
	if s != 2.5 || b != 3.5 || d != 2 {
		t.Error("Frequency Test Failed")
	}

}
