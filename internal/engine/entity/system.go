package entity

type System interface {
	Name()
	Update(entities []*entities)
}