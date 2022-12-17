package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerSetUp(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
