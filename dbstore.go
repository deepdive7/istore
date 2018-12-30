package istore

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"strings"
	"sync"
)

func NewLevelDBStore(dbPath string) (*LevelDBStore, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return nil, err
	}
	ls  := &LevelDBStore{
		db:db,
		keys: map[string]interface{}{},
		mu:&sync.RWMutex{},
	}
	return ls, nil
}

type LevelDBStore struct {
	db  *leveldb.DB
	keys map[string]interface{}
	opt *opt.Options
	wo  *opt.WriteOptions
	ro  *opt.ReadOptions
	mu  *sync.RWMutex
}

func (ls *LevelDBStore) Set(key []byte, val []byte) bool {
	ok, _ := ls.db.Has(key, ls.ro)
	if ok {
		return false
	}
	if  ls.db.Put(key, val, ls.wo) == nil {
		ls.keys[string(key)] = nil
		return true
	}
	return false
}

func (ls *LevelDBStore) Update(key []byte, val []byte) bool {
	if ls.db.Put(key, val, ls.wo) == nil{
		ls.keys[string(key)] = nil
		return true
	}
	return false
}

func (ls *LevelDBStore) Get(key []byte) ([]byte, error) {
	return ls.db.Get(key, ls.ro)
}

func (ls *LevelDBStore) Del(key []byte) {
	if ls.db.Delete(key, ls.wo) == nil {
		delete(ls.keys, string(key))
	}
}

func (ls *LevelDBStore) Search(keyPattern string) [][]byte {
	res := make([][]byte, 0, 0)
	for k := range ls.keys {
		if strings.Index(k, keyPattern) >= 0 {
			b, err := ls.Get([]byte(k))
			if err != nil {
				continue
			}
			res = append(res, b)
		}
	}
	return res
}

func (ls *LevelDBStore) Backup() bool {
	return false
}