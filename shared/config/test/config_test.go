package test

import (
	"log"
	"testing"
	cfg "tuzi-tiktok/config"
)

func TestConfig(t *testing.T) {
	log.Printf("%#v", cfg.DatabaseConfig)
}
