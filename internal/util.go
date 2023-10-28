package internal

import (
	"github.com/gofiber/fiber/v2"
)

func Startup(app *fiber.App) {
	if err := app.Listen("localhost:3000"); err != nil {
		Throw(4)
	}
}
