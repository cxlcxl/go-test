package app_cache

type Cache interface {
	Get(key string) (string, error)
	Set(key string, val ...interface{}) error
	Delete(key string) error
	Flush() error
}
