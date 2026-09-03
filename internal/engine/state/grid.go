package state

type Boundary struct {
	MinPosition Position `json:"minPosition"`
	MaxPosition Position `json:"maxPosition"`
}

type Grid struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Boundary Boundary `json:"boundary"`
	Tiles map[Position]*Tile `json:"tiles"`
}

func NewGrid(id, name, description string) *Grid

func (g *Grid) Contains(position Position) bool

func (g *Grid) CalculatePath(start Position, end Position) []Position

	