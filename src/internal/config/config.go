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
		port = "8085"
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://username:password@localhost:27017"
	}

	return &Config{
		Port:     port,
		MongoURI: mongoURI,
	}
}
