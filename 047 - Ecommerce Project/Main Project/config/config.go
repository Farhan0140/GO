package config

import (
	"fmt"
	"os"
	"strconv"
	"github.com/lpernett/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

var configuration Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to Load .env file", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is Required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is Required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is Required")
		os.Exit(1)
	}

	port, _ := strconv.Atoi(httpPort)

	configuration = Config {
		Version: version,
		ServiceName: serviceName,
		HttpPort: port,
	}
}

func GetConfig () Config {
	loadConfig()

	return configuration
}
