package workoutPkg

import (
	"testing"
)

//Testing GetMev
func TestGetMEV(t *testing.T) {
	input := []string{"ff", "Wmiddle", "Hmedium	", "Shigh", "Eadvanced", "A20s", "Dgood", "Sleepgood", "Stresslow", "HW4", "HRA4"}

	DefaultTraining(input)
}
