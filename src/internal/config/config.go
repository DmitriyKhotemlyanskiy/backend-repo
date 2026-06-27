package config

import "os"

type Config struct {
	Port     string
	MongoURI string
}

// LoadConfig reads configuration from environment variables
func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	return &Config{
		Port:     port,
		MongoURI: mongoURI,
	}
}
