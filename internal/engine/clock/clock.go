package clock

type Tick struct {
	Name string `json:"name"`
	Current uint64 `json:"current"`
	Max     uint64 `json:"max"`
}

type Clock struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Tick       Tick   `json:"tick"`
	TickCount  uint64 `json:"tickCount"`
}

func NewClock() *Clock

// NewTick will create a new tick with the given name and max.
func (c *Clock) NewTick(name string, max uint64)

// AddTicks will add the given number of ticks to the clock.
func (c *Clock) AddTicks(step uint64)

// Progress returns the progress of the clock as a float64.
func (c *Clock) Progress() float64
