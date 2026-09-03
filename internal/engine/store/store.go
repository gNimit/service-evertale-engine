package store

import "github.com/adrem/service-evertale-engine/internal/engine/state"

type Store interface {
	SaveWorld() error
	LoadWorld() (*state.World, error)
	
	SaveEntity() error
	LoadEntity() error

	SaveClock() error
	LoadClock() error
}

func NewStore() Store