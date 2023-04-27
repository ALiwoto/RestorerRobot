package entryManager

import (
	"sync"
)

func NewManager(triggers []rune) *EntryManager {
	return &EntryManager{
		entryMutex: &sync.RWMutex{},
		entryMap:   make(map[string]*entry),
		triggers:   triggers,
	}
}
