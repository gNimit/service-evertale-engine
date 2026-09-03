package entity

type System interface {
	Name() string
	Update(entities []*Entity)
}
