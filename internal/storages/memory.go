package storages

import "sync"

type MemoryStorage struct {
	sync.Mutex
	m map[string][]any
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func get() {

}
