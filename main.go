package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaeldajuda/tech-task-golang-api/handler"
)

func main() {

	app := fiber.New()

	app.Get("/api/login", handler.Login)
	app.Get("/api/register", handler.Register)
	app.Get("/api/tasks", handler.GetTasks)
	app.Get("/api/tasks/:id", handler.GetTask)
	app.Post("/api/tasks", handler.PostTask)
	app.Put("/api/tasks/:id", handler.PutTask)
	app.Delete("/api/tasks/:id", handler.DeleteTask)

	app.Listen(":3000")

}
