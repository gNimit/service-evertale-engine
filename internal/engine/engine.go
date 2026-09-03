package engine

import (
	"sync"

	"github.com/gNimit/service-evertale-engine/internal/engine/state"
)

type Engine struct {
	mu    sync.RWMutex
	World *state.World
}

func NewEngine() *Engine {
	return &Engine{
		mu:    sync.RWMutex{},
		World: nil,
	}
}

func (e *Engine) Run() error {
	// TODO: Implement run
	return nil
}
