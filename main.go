package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafaeldajuda/tech-task-golang-api/handler"
	"github.com/rafaeldajuda/tech-task-golang-api/utils"
)

func main() {

	// database
	db := utils.Connection()
	utils.ConfigTables(db)

	app := fiber.New()

	app.Post("/api/login", handler.PostLogin)
	app.Post("/api/register", handler.PostRegister)
	app.Get("/api/tasks", handler.GetTasks)
	app.Get("/api/tasks/:id", handler.GetTask)
	app.Post("/api/tasks", handler.PostTask)
	app.Put("/api/tasks/:id", handler.PutTask)
	app.Delete("/api/tasks/:id", handler.DeleteTask)

	app.Listen(":3000")

}
