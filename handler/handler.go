package handler

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/pkg"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) Handler {
	return Handler{db: db}
}

// authentication
func (h Handler) PostLogin(c *fiber.Ctx) (err error) {
	user := entity.User{}
	err = c.BodyParser(&user)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "3", Message: "login error"})
	}

	token, err := pkg.Login(user, h.db)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "4", Message: "login error"})
	}

	return c.Status(http.StatusOK).SendString(token)
}

func (h Handler) PostRegister(c *fiber.Ctx) (err error) {
	user := entity.User{}
	err = c.BodyParser(&user)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "1", Message: "register error"})
	}

	err = pkg.Register(user, h.db)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "2", Message: "register error"})
	}

	return c.Status(http.StatusCreated).SendString("")
}

// tasks
func (h Handler) GetTasks(c *fiber.Ctx) (err error) {
	return c.SendString("GetTasks")
}

func (h Handler) GetTask(c *fiber.Ctx) (err error) {
	return c.SendString("GetTask")
}

func (h Handler) PostTask(c *fiber.Ctx) (err error) {
	return c.SendString("PostTask")
}

func (h Handler) PutTask(c *fiber.Ctx) (err error) {
	return c.SendString("PutTask")
}

func (h Handler) DeleteTask(c *fiber.Ctx) (err error) {
	return c.SendString("DeleteTask")
}
