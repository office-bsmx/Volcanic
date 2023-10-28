package internal

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Startup(app *fiber.App) {
	addr := "localhost:3000"
	log.Println("Server is running on", addr)
	log.Fatal(app.Listen(addr))
}
