package settings

type Postgres struct {
	Host     string
	Port     string
	DB       int
	Password string
}

var PostgresSettings = &Postgres{}

func PostgresInit() {
	mapTo("postgres", PostgresSettings)
}
