package main

import "os"

type Config struct {
	ListenAddr string
	CloudURL   string
}

func LoadConfig() Config {
	return Config{
		ListenAddr: getEnv("LISTEN_ADDR", ":8080"),
		CloudURL:   getEnv("CLOUD_URL", "http://cloud-service:9000/events"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
