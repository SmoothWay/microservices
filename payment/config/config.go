package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetDSN() string {
	return getEnvironmentValue("DSN")
}

func GetApplicationPort() int {
	portstr := getEnvironmentValue("APP_PORT")
	port, err := strconv.Atoi(portstr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portstr)
	}

	return port
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment value is missing", key)
	}

	return os.Getenv(key)
}
