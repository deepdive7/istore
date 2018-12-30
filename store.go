package istore

type Store interface {
	Set(key []byte, val []byte) error
	Get(key []byte) ([]byte, error)
	Update(key []byte, val []byte) error
	Del(key []byte)
	Search(keyPattern string) [][]byte
	Backup() bool
}
