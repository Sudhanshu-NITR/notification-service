package main

import (
	"log"

	"github.com/Sudhanshu-NITR/notification-service/internal/server"
)

func main() {
	app := server.New()

	log.Println("🚀 Notification API starting on :8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
