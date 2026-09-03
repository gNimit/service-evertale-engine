package store

import (
	"context"

	"github.com/gNimit/service-evertale-engine/internal/engine/clock"
	"github.com/gNimit/service-evertale-engine/internal/engine/entity"
	"github.com/gNimit/service-evertale-engine/internal/engine/state"
)

type Store interface {
	SaveWorld(ctx context.Context, world *state.World) error
	LoadWorld(ctx context.Context, worldId string) (*state.World, error)

	SaveGrid(ctx context.Context, grid *state.Grid) error
	LoadGrid(ctx context.Context, gridId string) (*state.Grid, error)

	SaveEntity(ctx context.Context, entity *entity.Entity) error
	LoadEntity(ctx context.Context, entityId string) (*entity.Entity, error)

	SaveClock(ctx context.Context, clock *clock.Clock) error
	LoadClock(ctx context.Context, clockId string) (*clock.Clock, error)
}

func NewMemoryStore() Store
