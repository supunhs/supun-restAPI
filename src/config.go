package main

import (
	"os"
)

//MongoConfig - Structure for mongo config parameters
type MongoConfig struct {
	host     string
	port     string
	database string
	username string
	password string
}

var mongoConfig = MongoConfig{
	host:     getEnv("MONGO_HOST", "localhost"),
	port:     getEnv("MONGO_PORT", "27017"),
	database: getEnv("MONGO_DB", "airQuality"),
	username: getEnv("MONGO_USER", ""),
	password: getEnv("MONGO_PASSWORD", ""),
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
