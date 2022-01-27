package settings

type Redis struct {
	Host     string
	Port     string
	DB       int
	Password string
}

var RedisSettings = &Redis{}

func RedisInit() {
	mapTo("redis", RedisSettings)
}
