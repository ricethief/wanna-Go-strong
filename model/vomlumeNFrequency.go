package model

import (
	"fmt"
	"math"
)

var MevSquatStartPoint = []float64{7.5, 5.5, 4.5}    //starting point for squat mev
var MevBenchStartPoint = []float64{9, 8, 6.5}        //starting point for bench mev
var MevDeadliftStartPoint = []float64{5.5, 4.5, 2.5} //starting point for deadlift mev
var MrvSquatStartPoint = []float64{14, 9, 6}         //starting point for squat mrv
var MrvBenchStartPoint = []float64{17, 11, 8.5}      //starting point for bench mrv
var MrvDeadliftStartPoint = []float64{11, 7, 3.5}    //starting point for deadlift mrv
/*
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
|                             | Lower Volume |          |              |          |          |          | Higher Volume |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
|                             | -1.5         | -1       | -0.5         | 0        | +0.5     | +1       | +1.5          |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Gender                      |              |          |              | mm       | m        | f        | ff            |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Weight                      | super        |          | heavy        |          | middle   |          | light         |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Height                      |              | very     | tall         |          | medium   | short    |               |
|                             |              | tall     |              |          |          |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Strength                    |              |          | low          | moderate | high     | very     |               |
|                             |              |          |              |          |          | high     |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Experience                  |              | beginner | intermediate |          | advanced | very     |               |
|                             |              |          |              |          |          | advanced |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Age                         |              | <19      | 20s          | 30s      | 40s      | 50+      |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Diet                        |              |          | good         | avg      | poor     |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Sleep                       |              |          | good         | avg      | poor     |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Stress                      |              |          | low          | avg      | high     |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Historical Workload         |              | 1        | 2            | 3        | 4        | 5        |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Historical recovery ability |              | 1        | 2            | 3        | 4        | 5       |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
*/
var MEVCalcSheet = map[string]float64{
	"mm":            0,
	"m":             0.5,
	"f":             1,
	"ff":            1.5,
	"Wsuper":        -1.5,
	"Wheavy":        -0.5,
	"Wmiddle":       0.5,
	"Wlight":        1.5,
	"Hverytall":     -1,
	"Htall":         -0.5,
	"Hmedium":       0.5,
	"Hshort":        1,
	"Slow":          -0.5,
	"Smoderate":     0,
	"Shigh":         0.5,
	"Sveryhight":    1,
	"Ebeginner":     -1,
	"Eintermediate": -0.5,
	"Eadvanced":     0.5,
	"Everyadvanced": 1,
	"A<19":          -1,
	"A20s":          -0.5,
	"A30s":          0,
	"A40s":          0.5,
	"A50+":          1,
	"Dgood":         -0.5,
	"Davg":          0,
	"Dpoor":         0.5,
	"Sleepgood":     -0.5,
	"Sleepavg":      0,
	"Sleeppoor":     0.5,
	"Stresslow":     -0.5,
	"Stressavg":     0,
	"Stresshigh":    0.5,
	"HW1":           -1,
	"HW2":           -0.5,
	"HW3":           0,
	"HW4":           0.5,
	"HW5":           1,
	"HRA1":          -1,
	"HRA2":          -0.5,
	"HRA3":          0,
	"HRA4":          0.5,
	"HRA5":          1,
}

/*
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
|                             | Lower Volume |          |              |          |          |          | Higher Volume |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
|                             | -3           | -2       | -1           | 0        | +1       | +2       | +3            |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Gender                      |              |          |              | mm       | m        | f        | ff            |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Weight                      | super        |          | heavy        |          | middle   |          | light         |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Height                      |              | very     | tall         |          | medium   | short    |               |
|                             |              | tall     |              |          |          |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Strength                    |veryhigh      |          | high         | moderate | low      |          |               |
|                             |              |          |              |          |          |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Experience                  |              | very     |              |advanced  |          |intermediate              |
|                             |              | advanced |              |beginner  |          |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Age                         |50+           |          | 40s          | 30s      | 20s      | <19      |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Diet                        |              |poor      |              | avg      | good     |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Sleep                       |              |poor      |              | avg      | good     |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Stress                      |              |high      |              | avg      | low      |          |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Historical Workload         |              | 1        | 2            | 3        | 4        | 5        |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
| Historical recovery ability |              | 1        | 2            | 3        | 4        | 5       |               |
+-----------------------------+--------------+----------+--------------+----------+----------+----------+---------------+
*/
var MRVCalcSheet = map[string]float64{
	"mm":            0,
	"m":             1,
	"f":             2,
	"ff":            3,
	"Wsuper":        -3,
	"Wheavy":        -1,
	"Wmiddle":       1,
	"Wlight":        3,
	"Hverytall":     -2,
	"Htall":         -1,
	"Hmedium":       1,
	"Hshort":        2,
	"Slow":          1,
	"Smoderate":     0,
	"Shigh":         -1,
	"Sveryhight":    -3,
	"Eintermediate": 2,
	"Eadvanced":     0,
	"Ebeginner":     0,
	"Everyadvanced": -2,
	"A<19":          2,
	"A20s":          1,
	"A30s":          0,
	"A40s":          -1,
	"A50+":          -3,
	"Dgood":         1,
	"Davg":          0,
	"Dpoor":         -2,
	"Sleepgood":     1,
	"Sleepavg":      0,
	"Sleeppoor":     -2,
	"Stresslow":     1,
	"Stressavg":     0,
	"Stresshigh":    -2,
	"HW1":           -2,
	"HW2":           -1,
	"HW3":           0,
	"HW4":           1,
	"HW5":           2,
	"HRA1":          -2,
	"HRA2":          -1,
	"HRA3":          0,
	"HRA4":          1,
	"HRA5":          2,
}

//FrequencySheet
var FrequencySheet = map[string]float64{
	"mm":            11,
	"m":             14,
	"f":             17,
	"ff":            21,
	"Wsuper":        2,
	"Wheavy":        9,
	"Wmiddle":       17,
	"Wlight":        24,
	"Hverytall":     2,
	"Htall":         9,
	"Hmedium":       17,
	"Hshort":        24,
	"Slow":          24,
	"Smoderate":     17,
	"Shigh":         9,
	"Sveryhight":    2,
	"Eintermediate": 16,
	"Eadvanced":     12,
	"Ebeginner":     20,
	"Everyadvanced": 8,
	"A<19":          25,
	"A20s":          20,
	"A30s":          15,
	"A40s":          11,
	"A50+":          5,
	"Dgood":         16,
	"Davg":          11,
	"Dpoor":         4,
	"Sleepgood":     16,
	"Sleepavg":      11,
	"Sleeppoor":     4,
	"Stresslow":     16,
	"Stressavg":     11,
	"Stresshigh":    4,
	"HW1":           4,
	"HW2":           8,
	"HW3":           13,
	"HW4":           17,
	"HW5":           23,
	"HRA1":          4,
	"HRA2":          8,
	"HRA3":          13,
	"HRA4":          17,
	"HRA5":          23,
}

//make sure starting points has correct value
func reset() {

	MevSquatStartPoint = []float64{7.5, 5.5, 4.5}    //starting point for squat mev
	MevBenchStartPoint = []float64{9, 8, 6.5}        //starting point for bench mev
	MevDeadliftStartPoint = []float64{5.5, 4.5, 2.5} //starting point for deadlift mev
	MrvSquatStartPoint = []float64{14, 9, 6}         //starting point for squat mrv
	MrvBenchStartPoint = []float64{17, 11, 8.5}      //starting point for bench mrv
	MrvDeadliftStartPoint = []float64{11, 7, 3.5}    //starting point for deadlift mrv

}

/*
GetMEV will calculate user's MEV (Maximum Efective Volume)
Return squat[]{hypertrophy_MEV int, strength_MEV int, peaking_MEV int}

reference: Juggernaut Training Systems-Program Design Manual p35.
0.5 value get rounded up in MEV
*/
func GetMEV(userinfo []string) (squatmev, benchmev, deadliftmev []float64) {
	Change := 0.00 //start at 0
	for _, s := range userinfo {
		Change = Change + MEVCalcSheet[s] //getting and adding it's value from calcsheet
		fmt.Println("key :", s, "value :", MEVCalcSheet[s])
	}
	fmt.Println("Change : ", Change)
	s := []float64{MevSquatStartPoint[0] + Change,
		MevSquatStartPoint[1] + Change,
		MevSquatStartPoint[2] + Change,
	}
	b := []float64{MevBenchStartPoint[0] + Change,
		MevBenchStartPoint[1] + Change,
		MevBenchStartPoint[2] + Change,
	}
	d := []float64{MevDeadliftStartPoint[0] + Change,
		MevDeadliftStartPoint[1] + Change,
		MevDeadliftStartPoint[2] + Change,
	}
	return s, b, d
}

/*
GetMRV will calculate user's MRV (Maximum Recoverable Volume)
Return squat[]{hypertrophy_MRV int, strength_MRV int, peaking_MRV int}

reference: Juggernaut Training Systems-Program Design Manual p35.
*/
func GetMRV(userinfo []string) (squatmev, benchmev, deadliftmev []float64) {
	Change := 0.00 //start at 0
	for _, s := range userinfo {
		Change = Change + MRVCalcSheet[s] //getting and adding it's value from calcsheet
		fmt.Println("key :", s, "value :", MRVCalcSheet[s])
	}
	fmt.Println("Change : ", Change)
	s := []float64{MrvSquatStartPoint[0] + Change,
		MrvSquatStartPoint[1] + Change,
		MrvSquatStartPoint[2] + Change,
	}
	b := []float64{MrvBenchStartPoint[0] + Change,
		MrvBenchStartPoint[1] + Change,
		MrvBenchStartPoint[2] + Change,
	}
	d := []float64{MrvDeadliftStartPoint[0] + Change,
		MrvDeadliftStartPoint[1] + Change,
		MrvDeadliftStartPoint[2] + Change,
	}
	return s, b, d
}

//DeterminingFrequency
func GetFreqiency(userinfo []string) (sfrequency, bfrequency, dfreqiency float64) {
	change := 0.00 //start at 0
	for _, s := range userinfo {
		change = change + FrequencySheet[s] //getting and adding it's value from calcsheet
		fmt.Println("fkey :", s, "value :", FrequencySheet[s])
	}
	//caclculate the avg of the info
	change = math.Round((change/float64(len(userinfo)) + 2))

	fmt.Println("Change for frequency : ", change)
	if change >= 5 && change < 10 {
		sfrequency, bfrequency, dfreqiency = 1, 2, 1
	} else if change >= 10 && change < 13 {
		sfrequency, bfrequency, dfreqiency = 1.5, 2.5, 1.5
	} else if change >= 13 && change < 15 {
		sfrequency, bfrequency, dfreqiency = 2, 3, 1.5
	} else if change >= 15 && change < 17 {
		sfrequency, bfrequency, dfreqiency = 2, 3, 2
	} else if change >= 17 && change < 19 {
		sfrequency, bfrequency, dfreqiency = 2.5, 3.5, 2
	} else if change >= 19 && change < 25 {
		sfrequency, bfrequency, dfreqiency = 3, 4, 2.5
	} else {
		sfrequency, bfrequency, dfreqiency = 3.5, 4.5, 3
	}
	return sfrequency, bfrequency, dfreqiency
}

// DefaultTraining will generate basic strenth training
// Parameters sex int,height float64,weight float64,exp int,lenOfWeek int
// return training
func DefaultTraining(input []string) {
	s, b, d := GetMEV(input)
	fmt.Println("[sets per week]Mev for Squat:", s, "Benchpress:", b, "Deadlift:", d)
	s1, b1, d1 := GetMRV(input)
	fmt.Println("[sets per week]Mrv for Squat:", s1, "Benchpress:", b1, "Deadlift:", d1)
	ss, bb, dd := GetFreqiency(input)
	fmt.Println("[session per week]frequency for Squat:", ss, "Benchpress:", bb, "Deadlift:", dd)
}
