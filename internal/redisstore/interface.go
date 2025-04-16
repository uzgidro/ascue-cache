package redisstore

type Store interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}
