package enteryManager

import (
	"sync"
)

func NewManager(triggers []rune) *EnteryManager {
	return &EnteryManager{
		enteryMutex: &sync.Mutex{},
		enteryMap:   make(map[string]*Entery),
		triggers:    triggers,
	}
}
