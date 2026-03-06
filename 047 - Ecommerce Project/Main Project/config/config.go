package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type DBConfig struct {
	Host           string
	Port           int
	DBName         string
	User           string
	Password       string
	EnabledSSLMode bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	SecretKey   string
	DB          *DBConfig
}

var configuration *Config

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

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("JWT Secret Key is required")
		os.Exit(1)
	}

	// DataBase Config
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("Database Host is required")
		os.Exit(1)
	}

	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		fmt.Println("Invalid Type Data Base Port Type")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("Database Name is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("Database User is required")
		os.Exit(1)
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		fmt.Println("Database Password is required")
		os.Exit(1)
	}

	enabledSslModeStr := os.Getenv("DB_ENABLED_SSL_MODE")
	enabledSslMode, err := strconv.ParseBool(enabledSslModeStr)
	if err != nil {
		fmt.Println("Invalid Type Data Base SSL Type")
	}

	dbConfig := &DBConfig{
		Host:           dbHost,
		Port:           dbPort,
		DBName:         dbName,
		User:           dbUser,
		Password:       dbPass,
		EnabledSSLMode: enabledSslMode,
	}

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
		SecretKey:   jwtSecretKey,
		DB:          dbConfig,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}

	return configuration
}
