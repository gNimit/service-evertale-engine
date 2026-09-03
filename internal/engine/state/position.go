package state

type Position struct {
	X       int  `json:"x"`
	Y       int  `json:"y"`
	IsFixed bool `json:"isFixed"`
}

func NewPosition(x, y, z int) *Position