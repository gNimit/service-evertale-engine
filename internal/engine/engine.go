package engine

import (
	"sync"

	"github.com/adrem/service-evertale-engine/internal/engine/state"
)

type Command struct {

}

type Engine struct {
	mu sync.RWMutex
	World *state.World
}

func NewEngine() *Engine