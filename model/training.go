package model

type Training struct {
	Name  string
	Weeks []Week
}

func NewTraining(trainingName string) Training {
	t := Training{
		Name: trainingName,
	}
	return t
}

//addWeek
func (t *Training) AddWeek(w Week) {
	t.Weeks = append(t.Weeks, w)
}

//removeWeek
func (t *Training) RemoveWeek(w Week) {
	t.Weeks = t.Weeks[:len(t.Weeks)-1]
}
