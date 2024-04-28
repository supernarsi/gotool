package gotool

import "sync"

type SyncMap struct {
	mu     sync.RWMutex
	values map[any]struct{}
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		values: make(map[any]struct{}),
	}
}

func (s *SyncMap) IsExist(key any) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.values[key]; ok {
		return true
	}
	return false
}

func (s *SyncMap) Lock(key any) {
	s.mu.Lock()
	s.values[key] = struct{}{}
	s.mu.Unlock()
}

func (s *SyncMap) Unlock(key any) {
	s.mu.Lock()
	delete(s.values, key)
	s.mu.Unlock()
}
