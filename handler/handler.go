package handler

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/pkg"
	"github.com/rafaeldajuda/tech-task-golang-api/utils"
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

	// gerar token
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
	// validar token
	token := c.GetReqHeaders()["Token"][0]
	_, _, err = utils.ValidToken(token)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "5", Message: "get tasks error"})
	}
	return c.SendString("GetTasks")
}

func (h Handler) GetTask(c *fiber.Ctx) (err error) {
	return c.SendString("GetTask")
}

func (h Handler) PostTask(c *fiber.Ctx) (err error) {
	task := entity.Task{}
	err = c.BodyParser(&task)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "8", Message: "post task error"})
	}

	// validar token
	if len(c.GetReqHeaders()["Token"]) == 0 {
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "6", Message: "post task error"})
	}
	token := c.GetReqHeaders()["Token"][0]
	id, email, err := utils.ValidToken(token)
	if err != nil {
		log.Errorf("error: %s", err.Error())
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "7", Message: "post task error"})
	}

	// guardar task - validar dados usuario
	pkg.PostTask(id, email, task)

	return c.Status(http.StatusCreated).SendString("")
}

func (h Handler) PutTask(c *fiber.Ctx) (err error) {
	return c.SendString("PutTask")
}

func (h Handler) DeleteTask(c *fiber.Ctx) (err error) {
	return c.SendString("DeleteTask")
}
