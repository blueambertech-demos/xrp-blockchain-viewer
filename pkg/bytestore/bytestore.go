package bytestore

import "sync"

type Store struct {
	data sync.Map
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) SetData(key string, data []byte) {
	s.data.Store(key, data)
}

func (s *Store) GetData(key string) ([]byte, bool) {
	d, ok := s.data.Load(key)
	if !ok {
		return nil, false
	}
	return d.([]byte), true
}
