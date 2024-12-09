package main

import (
	"gateway_service/internal/config"
	"gateway_service/internal/routing"
)

func main() {
	cfg := config.LoadConfig()
	r := routing.SetupRouter(cfg)
	r.Run(":8080") // gateway слушает на порту 8080
}
