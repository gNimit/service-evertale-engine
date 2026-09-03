package entity

type ComponentType string

const (
	CompStat ComponentType = "stat"
	CompInventory ComponentType = "inventory"
	CompTraits ComponentType = "traits"
	CompAiContext ComponentType = "aiContext"
)

type Component interface {
	Type() ComponentType
}

type StatsComponent struct {
	Values 	map[string]int `json:"values"`	
}

func (s *StatsComponent) Type() ComponentType

func (s *StatsComponent) Get(stat string) int

func (s *StatsComponent) Set(stat string, value int)

type InventoryComponent struct

func (i *InventoryComponent) Type() ComponentType

func (i *InventoryComponent) Get(item string) int

func (i *InventoryComponent) Add(item string, count int)

func (i *InventoryComponent) Remove(item string, count int) bool

type TraitsComponent struct {
	Traits []string `json:"traits"`
}

func (t *TraitsComponent) Type() ComponentType

func (t *TraitsComponent) Has(trait string) bool

func (t *TraitsComponent) Add(trait string)

func (t *TraitsComponent) Remove(trait string) bool

type AiContextComponent struct {
	Memory []string `json:"memory"`
}

func (a *AiContextComponent) Type() ComponentType

func (a *AiContextComponent) Add(memory string)
