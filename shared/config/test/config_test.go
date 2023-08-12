package test

import (
	"log"
	"testing"
	cfg "tuzi-tiktok/config"
)

func TestConfig(t *testing.T) {
	log.Printf(cfg.LoggerConfig.Encoding)
	log.Printf("%#v", cfg.LoggerConfig)
}
