package config

type Config struct {
	Database struct {
		DSN string
	}
	Server struct {
		Port string
	}
}

func LoadConfig() (*Config, error) {
	// Load from .env, file, etc.
	return &Config{
		Database: struct{ DSN string }{
			DSN: "host=localhost user=postgres password=8231 dbname=gorm_boom port=5432 sslmode=disable TimeZone=Asia/Bangkok",
		},
		Server: struct{ Port string }{
			Port: "3001",
		},
	}, nil
}
