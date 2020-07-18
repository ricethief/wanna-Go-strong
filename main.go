package main

import (
	"fmt"
	"wanna-Go-strong/model"
)

func main() {
	fmt.Println("Hello world")
	e := model.NewExercise("Squat", "")
	e.AddSet(5, 5)
	d := model.NewDay("Lag day")
	p := &d
	p.AddDay(e)

	fmt.Println("*******************")
	fmt.Println("date:{}", d)
	fmt.Println("entered obj:{}", e)
}
