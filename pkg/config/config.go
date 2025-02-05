package beep

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func InitializeConfig[C any](path string) (C, error) {
	var config C
	err := godotenv.Load(path)
	if err != nil {
		return config, err
	}
	err = env.Parse(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
