package cache

//Cache interface
type Cache interface {
	GetJSON(key string, reply interface{}) error
	GetString(key string) (string, error)
	SetJSON(key string, val interface{}, timeout int64) error
	SetString(key, val string, timeout int64) error
	IsExist(key string) bool
	Delete(key string) error
}
