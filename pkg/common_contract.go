package pkg

import "github.com/gofiber/fiber/v2"

type Container interface {
	HandleApi(app fiber.Router)
	HandleWeb(app fiber.Router)
}
