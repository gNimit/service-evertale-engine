package state

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
	TargetTileId string `json:"targetTileId"`
	TargetTilePosition *Position `json:"targetTilePosition"`
	TravelCost uint64 `json:"travelCost"`
}

type Tile struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Glyph rune `json:"glyph"`
	Color string `json:"color"`
	Position Position `json:"position"`
	EntityIds []string `json:"entityIds"`
	Edges map[Direction]*Edge `json:"edges"`
}

func NewTile(id, name, description string, position Position) *Tile

func (t *Tile) AddEntity(entityId string) error

func (t *Tile) RemoveEntity(entityId string) error



