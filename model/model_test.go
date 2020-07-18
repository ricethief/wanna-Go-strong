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

//testing
