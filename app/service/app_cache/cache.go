package app_cache

type AppCache struct {
	Driver string
}

func NewCache(driver string) *AppCache {
	if driver == "" {
		driver = "redis"
	}
	return &AppCache{
		Driver: driver,
	}
}

func setDriver() {

}
