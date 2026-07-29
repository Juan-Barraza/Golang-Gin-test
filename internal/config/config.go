package config

import "os"

type Config struct {
	Port           string
	DataFilePath   string
	BoostrapApiUrl string
}

func Load() *Config {
	return &Config{
		Port:           os.Getenv("PORT"),
		DataFilePath:   os.Getenv("DATA_FILE_PATH"),
		BoostrapApiUrl: os.Getenv("BOOSTR_API_URL"),
	}
}
