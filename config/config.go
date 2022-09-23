package config

import (
	"time"
)

type Config struct {
	scrapeList         []string
	timeBetweenScrapes time.Duration `default:1000`
}

func New() Config {
	var cfg Config

	return cfg
}
