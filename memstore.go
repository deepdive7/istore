package istore

import (
	"errors"
	"strings"
	"sync"
)

func NewMemStore(prefix string) *MemStore {
	return &MemStore{prefix: prefix}
}

type MemStore struct {
	store  map[string][]byte
	mux    *sync.RWMutex
	prefix string
}

func (ms *MemStore) Set(key []byte, val []byte) bool {
	ms.mux.Lock()
	defer ms.mux.Unlock()
	k := string(key)
	if _, ok := ms.store[k]; ok {
		return false
	}
	ms.store[k] = val
	return true
}

func (ms *MemStore) Update(key []byte, val []byte) bool {
	ms.mux.Lock()
	defer ms.mux.Unlock()
	k := string(key)
	if _, ok := ms.store[k]; ok {
		ms.store[k] = val
		return true
	}
	return false
}

func (ms *MemStore) Get(key []byte) ([]byte, error) {
	ms.mux.RLock()
	defer ms.mux.RUnlock()
	k := string(key)
	if _, ok := ms.store[k]; ok {
		return ms.store[k], nil
	}
	return nil, errors.New("no suck key " + string(k))
}

func (ms *MemStore) Del(key []byte) {
	ms.mux.Lock()
	defer ms.mux.Unlock()
	delete(ms.store, string(key))
}

func (ms *MemStore) Search(keyPattern string) ([][]byte) {
	res := make([][]byte, 0, 0)
	for k, v := range ms.store {
		if strings.Index(k, keyPattern) >= 0 {
			res = append(res, v)
		}
	}
	return res
}

func (ms *MemStore) Backup() bool{
	return false
}
