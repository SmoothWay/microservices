package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironment("ENV")
}

func GetDSN() string {
	return getEnvironment("DSN")
}

func GetApplicationPort() int {
	portStr := getEnvironment("APP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", err)
	}

	return port
}

func GetPaymentServiceUrl() string {
	return getEnvironment("PAYMENT_SERVICE_URL")
}

func getEnvironment(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing", key)
	}

	return os.Getenv(key)
}
