package main

import (
	"log"

	"github.com/ZanzyTHEbar/goflexpro/internal/server"
	"github.com/ZanzyTHEbar/goflexpro/pkgs/config"
)

func main() {
	config, err := config.NewConfig(nil, nil)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server.
		NewServer(config).
		Start()
}
