package pkg

import "github.com/gofiber/fiber/v2"

type Router interface {
	Setup()
}

type router struct {
	app        *fiber.App
	containers []Container
}

func NewRouter(app *fiber.App, c []Container) Router {
	return &router{
		app:        app,
		containers: c,
	}
}

func (r router) Setup() {
	api := r.app.Group("/api")
	for _, container := range r.containers {
		container.HandleApi(api)
	}

	web := r.app.Group("/")
	for _, container := range r.containers {
		container.HandleWeb(web)
	}
}
