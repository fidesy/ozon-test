package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string   `yaml:"host"`
	Port     string   `yaml:"port"`
	GrpcPort string   `yaml:"grpc_port"`
	Database string   `yaml:"database"`
	Postgres Postgres `yaml:"postgres"`
}

type Postgres struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

var Default = Config{
	Host:     "http://localhost",
	Port:     "8000",
	GrpcPort: "5001",
	Database: "postgres",
	Postgres: Postgres{
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     "5432",
		DBName:   "postgres",
		SSLMode:  "disable",
	},
}

func Load(filepath string) (Config, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}

	var conf Config
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
