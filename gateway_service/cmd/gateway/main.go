package main

import (
	"gateway_service/internal/config"
	"gateway_service/internal/routing"
)

func main() {
	cfg := config.LoadConfig()
	r := routing.SetupRouter(cfg)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
