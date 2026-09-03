package main

import (
	"log"

	"github.com/gNimit/service-evertale-engine/internal/engine"
)

func main() {
	engine := engine.NewEngine()
	if err := engine.Run(); err != nil {
		log.Fatalf("failed to run engine: %v", err)
	}
}
