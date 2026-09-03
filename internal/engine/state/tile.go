package state

import "errors"

type Direction string

const (
	DirectionNorth     Direction = "north"
	DirectionEast      Direction = "east"
	DirectionSouth     Direction = "south"
	DirectionWest      Direction = "west"
	DirectionNorthEast Direction = "northeast"
	DirectionNorthWest Direction = "northwest"
	DirectionSouthEast Direction = "southeast"
	DirectionSouthWest Direction = "southwest"
)

type Edge struct {
	TargetTileId       string    `json:"targetTileId"`
	TargetTilePosition *Position `json:"targetTilePosition"`
	TravelCost         uint64    `json:"travelCost"`
}

type Tile struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Glyph       rune                `json:"glyph"`
	Color       string              `json:"color"`
	Position    Position            `json:"position"`
	EntityIds   map[string]bool     `json:"entityIds"`
	Edges       map[Direction]*Edge `json:"edges"`
}

func NewTile(id, name, description string, position Position) *Tile {
	return &Tile{
		ID:          id,
		Name:        name,
		Description: description,
		Glyph:       ' ',
		Color:       "#000000",
		Position:    position,
		EntityIds:   make(map[string]bool),
		Edges:       make(map[Direction]*Edge),
	}
}

func (t *Tile) AddEntity(entityId string) error {
	if t.EntityIds == nil {
		t.EntityIds = make(map[string]bool)
	}

	if t.EntityIds[entityId] {
		return errors.New("entity already exists")
	}
	t.EntityIds[entityId] = true
	return nil
}

func (t *Tile) RemoveEntity(entityId string) error {
	if !t.EntityIds[entityId] {
		return errors.New("entity does not exist")
	}
	delete(t.EntityIds, entityId)
	return nil
}
