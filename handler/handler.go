package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/pkg"
)

// authentication
func PostLogin(c *fiber.Ctx) (err error) {
	return c.SendString("Login")
}

func PostRegister(c *fiber.Ctx) (err error) {
	user := entity.User{}
	err = c.BodyParser(&user)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "1", Message: "register error"})
	}
	err = pkg.Register(user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "2", Message: "register error"})
	}
	return c.Status(http.StatusCreated).SendString("")
}

// tasks
func GetTasks(c *fiber.Ctx) (err error) {
	return c.SendString("GetTasks")
}

func GetTask(c *fiber.Ctx) (err error) {
	return c.SendString("GetTask")
}

func PostTask(c *fiber.Ctx) (err error) {
	return c.SendString("PostTask")
}

func PutTask(c *fiber.Ctx) (err error) {
	return c.SendString("PutTask")
}

func DeleteTask(c *fiber.Ctx) (err error) {
	return c.SendString("DeleteTask")
}
