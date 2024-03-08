package handler

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) (err error) {
	return c.SendString("Login")
}

func Register(c *fiber.Ctx) (err error) {
	return c.SendString("Register")
}

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
