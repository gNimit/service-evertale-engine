package state

type Boundary struct {
	MinPosition Position `json:"minPosition"`
	MaxPosition Position `json:"maxPosition"`
}

type Grid struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Boundary    Boundary           `json:"boundary"`
	Tiles       map[Position]*Tile `json:"tiles"`
}

func NewGrid(id, name, description string) *Grid {
	return &Grid{
		ID:          id,
		Name:        name,
		Description: description,
		Boundary:    Boundary{},
		Tiles:       make(map[Position]*Tile),
	}
}

func (g *Grid) Contains(position Position) bool {
	return position.X >= g.Boundary.MinPosition.X && position.X <= g.Boundary.MaxPosition.X &&
		position.Y >= g.Boundary.MinPosition.Y && position.Y <= g.Boundary.MaxPosition.Y
}

func (g *Grid) CalculatePath(start Position, end Position) []Position {
	// TODO: Implement pathfinding
	return nil
}
