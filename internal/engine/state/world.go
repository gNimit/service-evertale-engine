package state

type World struct {
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	LoadedGrids   map[string]*Grid `json:"loadedGrids"`
	CurrentGridId string           `json:"currentGridId"`
}

func NewWorld(id, name, description string) *World {
	return &World{
		ID:            id,
		Name:          name,
		Description:   description,
		LoadedGrids:   make(map[string]*Grid),
		CurrentGridId: "",
	}
}
