package main

import (
	"os"
	"strconv"
)

type Config struct {
	grpcPort     int
	httpPort     int
	dbConnection string
}

func ParseEnv() *Config {
	grpcPort, err := strconv.ParseInt(os.Getenv("GRPC_PORT"), 10, 0)
	if err != nil {
		grpcPort = 10000
	}

	httpPort, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 0)
	if err != nil {
		httpPort = 11000
	}

	dbConnection := os.Getenv("DB_CONNECTION")

	return &Config{
		grpcPort:     int(grpcPort),
		httpPort:     int(httpPort),
		dbConnection: dbConnection,
	}
}
