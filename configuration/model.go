package configuration

type Config struct {
	Port     string `json:"port"`
	Database DB     `json:"database"`
}

type DB struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
	Port         string `json:"port"`
	Host         string `json:"host"`
}
