package state

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewPosition(x, y int, isFixed bool) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}
