package entity

import (
	"sync"

	"github.com/adrem/service-evertale-engine/internal/engine/state"
)

type Entity struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Glyph rune
	Position state.Position
	Description string `json:"description"`
	Tags map[string]bool `json:"tags"`
	Components map[string]*Component `json:"components"`
}

func NewEntity(name, description string, position state.Position) *Entity

func (e *Entity) SetPosition(position state.Position)

func (e *Entity) SetGlyph(glyph rune)

func (e *Entity) HasTag(tag string) bool

func (e *Entity) AddTag(tag string) bool

func (e *Entity) RemoveTag(tag string) bool

func (e *Entity) HasComponent(component string) bool

func (e *Entity) AddComponent(component *Component) bool

func (e *Entity) RemoveComponent(component string) bool


