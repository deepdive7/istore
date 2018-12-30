package istore

type Store interface {
	Set(key []byte, val []byte) bool
	Update(key []byte, val []byte) bool
	Get(key []byte) ([]byte, error)
	Del(key []byte)
	Search(keyPattern string) [][]byte
	Backup() bool
}
