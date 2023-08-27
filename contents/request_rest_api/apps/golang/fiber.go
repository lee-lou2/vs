package golang

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func MainFiber() {
	app := fiber.New()

	app.Get("/health/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
