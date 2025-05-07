package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	Redis struct {
		Addr     string `json:"addr"`
		DB       int    `json:"db"`
		Password string `json:"password"`
	} `json:"redis"`
	CSVPath        string `json:"csv_path"`
	OutputCSVPath  string `json:"output_csv_path"`
	MaxRowsToRead  int    `json:"max_rows_to_read"`
	KeyCachePrefix string `json:"key_cache_prefix"`
}

func LoadConfig(path string) (*AppConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg AppConfig
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
