package bytestore

import (
	"context"
	"sync"
	"time"
)

type Store struct {
	mutex sync.RWMutex
	data  map[string]storedData
}

type storedData struct {
	data   []byte
	expiry int64
}

const monitorRefreshSecs = 1

// NewStore creates a new memory store to hold byte slices
func NewStore() *Store {
	return &Store{
		data: map[string]storedData{},
	}
}

// SetData sets data in the store using the provided key. A goroutine is launched to monitor the expiry of this data, when the expiry time
// is reached the data will be deleted
func (s *Store) SetData(ctx context.Context, key string, data []byte, expiry int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	sd := storedData{
		data:   data,
		expiry: expiry,
	}
	var exists bool
	if _, exists = s.data[key]; !exists {
		s.data[key] = sd
		go monitorKeyExpiry(ctx, s, key)
	} else {
		s.data[key] = sd
	}
}

// GetData retrieves data stored against a key, if no data is stored against this key it will return a nil byte array and the
// bool returned will be false
func (s *Store) GetData(key string) ([]byte, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	d, ok := s.data[key]
	if !ok {
		return nil, false
	}
	return d.data, true
}

// DeleteData removes an entry from the store with the given key
func (s *Store) DeleteData(key string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, key)
}

func monitorKeyExpiry(ctx context.Context, store *Store, key string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			exit := func() bool {
				store.mutex.Lock()
				defer store.mutex.Unlock()
				d, ok := store.data[key]
				if !ok {
					return true
				}
				if time.Now().Unix() > d.expiry {
					delete(store.data, key)
					return true
				}
				return false
			}()
			if exit {
				return
			}
			time.Sleep(time.Second * monitorRefreshSecs)
		}
	}
}
