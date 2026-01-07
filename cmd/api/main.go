package main

import (
	"log"

	"github.com/Sudhanshu-NITR/notification-service/internal/config"
	"github.com/Sudhanshu-NITR/notification-service/internal/server"
)

func main() {
	cfg := config.Load()

	app := server.New(cfg)

	log.Printf("🚀 %s running on :%s (%s)",
		cfg.App.Name,
		cfg.App.Port,
		cfg.App.Env,
	)

	if err := app.Run(":" + cfg.App.Port); err != nil {
		log.Fatal(err)
	}
}
