package config

type Config struct {
	Database Database
	Port     uint
}

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
}
