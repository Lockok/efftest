package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP HTTPConfig
	DB   DBConfig
}

type HTTPConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Load() (*Config, error) {

	_ = godotenv.Load()

	httpPort, err := requireEnv("HTTP_PORT")
	if err != nil {
		return nil, err
	}

	dbHost, err := requireEnv("DB_HOST")
	if err != nil {
		return nil, err
	}

	dbPort, err := requireEnv("DB_PORT")
	if err != nil {
		return nil, err
	}

	dbUser, err := requireEnv("DB_USER")
	if err != nil {
		return nil, err
	}

	dbPassword, err := requireEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	dbName, err := requireEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	
	cfg := &Config{
		HTTP: HTTPConfig{
			Port: httpPort,
		},
		DB: DBConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:    dbUser,
			Password: dbPassword,
			Name:     dbName,
		},
	}

	return cfg, nil

}

func requireEnv(key string) (string, error) {
    value := os.Getenv(key)
    if value == "" {
        return "", fmt.Errorf("%s is not set", key)
    }
    return value, nil
}

func (db DBConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Name)
}