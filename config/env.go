package config

import (
	"log"
	"os"
	"strconv"
)

var Env config = initConfig()

type config struct {
	Port       int
	DBType     string
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

func initConfig() config {
	return config{
		Port:       getEnvAsInt("APP_PORT", 3000),
		DBType:     getEnv("DB_TYPE", "postgres"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "mydb"),
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_USERNAME", "root"),
	}
}

func getEnv(name, fallback string) string {
	value, ok := os.LookupEnv(name)
	if ok {
		return value
	}
	return fallback
}

func getEnvAsInt(name string, fallback int) int {
	value, ok := os.LookupEnv(name)
	if ok {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Println(err)
			return fallback
		}
		return v
	}
	return fallback
}
