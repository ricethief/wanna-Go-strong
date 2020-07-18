package model

type Week struct {
	BlockType string
	Days      map[string]Day
}

func NewWeek(blockType string) Week {
	DList := make(map[string]Day)
	w := Week{
		BlockType: blockType,
		Days:      DList,
	}
	return w
}

func (w *Week) AddDay(date string, d Day) {
	w.Days[date] = d

}
