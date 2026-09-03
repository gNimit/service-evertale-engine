package event

import "sync"

type Event struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func NewEvent(name, description string) *Event
