package workoutPkg

import (
	"fmt"
	"testing"
)

//Testing GetMev
func TestGetMEV(t *testing.T) {
	input := []string{"ff", "Wmiddle", "Hmedium", "Smoderate", "Eintermediate",
		"A20s", "Davg", "Sleepavg", "Stresslow", "HW2", "HRA2"}

	DefaultTraining(input)
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
