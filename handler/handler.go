package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	// input log
	utils.InputLog(string(c.Context().Path()), c.Request().Header.String(), string(c.Body()))

	user := entity.User{}
	err = c.BodyParser(&user)
	if err != nil {
		bodyError := entity.AppError{Code: "3", Message: "login error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// gerar token
	token, err := pkg.Login(user, h.db)
	if err != nil {
		bodyError := entity.AppError{Code: "3", Message: "login error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// response
	utils.ResponseLog(token, http.StatusOK)
	return c.Status(http.StatusOK).SendString(token)
}

func (h Handler) PostRegister(c *fiber.Ctx) (err error) {
	// input log
	utils.InputLog(string(c.Context().Path()), c.Request().Header.String(), string(c.Body()))

	user := entity.User{}
	err = c.BodyParser(&user)
	if err != nil {
		bodyError := entity.AppError{Code: "1", Message: "register error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	err = pkg.Register(user, h.db)
	if err != nil {
		bodyError := entity.AppError{Code: "2", Message: "register error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// response
	utils.ResponseLog("", http.StatusCreated)
	return c.Status(http.StatusCreated).SendString("")
}

// tasks
func (h Handler) GetTasks(c *fiber.Ctx) (err error) {
	// input log
	utils.InputLog(string(c.Context().Path()), c.Request().Header.String(), string(c.Body()))

	// validar token
	if len(c.GetReqHeaders()["Token"]) == 0 {
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "6", Message: "get all tasks error"})
	}
	token := c.GetReqHeaders()["Token"][0]
	id, email, err := utils.ValidToken(token)
	if err != nil {
		bodyError := entity.AppError{Code: "7", Message: "get all tasks error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// listar tasks
	tasks, err := utils.SelectTasks(0, id, email, h.db)
	if err != nil {
		bodyError := entity.AppError{Code: "8", Message: "get all tasks error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// response
	responseBody, _ := json.Marshal(tasks)
	utils.ResponseLog(string(responseBody), http.StatusOK)
	return c.Status(http.StatusOK).JSON(tasks)
}

func (h Handler) GetTask(c *fiber.Ctx) (err error) {
	// input log
	utils.InputLog(string(c.Context().Path()), c.Request().Header.String(), string(c.Body()))

	// validar token
	if len(c.GetReqHeaders()["Token"]) == 0 {
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "6", Message: "get task error"})
	}
	token := c.GetReqHeaders()["Token"][0]
	id, email, err := utils.ValidToken(token)
	if err != nil {
		bodyError := entity.AppError{Code: "7", Message: "get task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// pegar id task
	idTask, err := c.ParamsInt("id")
	if err != nil {
		if err != nil {
			bodyError := entity.AppError{Code: "9", Message: "get task error"}
			bodyErrorS, _ := json.Marshal(bodyError)
			utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
			return c.Status(http.StatusBadRequest).JSON(bodyError)
		}
	}

	// listar tasks
	task, err := pkg.GetTask(int64(idTask), id, email, h.db)
	if err != nil {
		bodyError := entity.AppError{Code: "8", Message: "get task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}
	if task.ID == 0 {
		// response
		utils.ResponseLog("null", http.StatusOK)
		return c.Status(http.StatusOK).JSON(nil)
	}

	// response
	responseBody, _ := json.Marshal(task)
	utils.ResponseLog(string(responseBody), http.StatusOK)
	return c.Status(http.StatusOK).JSON(task)
}

func (h Handler) PostTask(c *fiber.Ctx) (err error) {
	// input log
	utils.InputLog(string(c.Context().Path()), c.Request().Header.String(), string(c.Body()))

	// validar token
	if len(c.GetReqHeaders()["Token"]) == 0 {
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "6", Message: "post task error"})
	}
	token := c.GetReqHeaders()["Token"][0]
	id, email, err := utils.ValidToken(token)
	if err != nil {
		bodyError := entity.AppError{Code: "7", Message: "post task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// pegar body
	task := entity.Task{}
	err = c.BodyParser(&task)
	if err != nil {
		bodyError := entity.AppError{Code: "8", Message: "post task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// guardar task - validar dados usuario
	idTask, err := pkg.PostTask(id, email, task, h.db)
	if err != nil {
		bodyError := entity.AppError{Code: "9", Message: "post task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// response
	response := entity.PostTaskSuccess{IDTask: idTask}
	responseBody, _ := json.Marshal(response)
	utils.ResponseLog(string(responseBody), http.StatusOK)
	return c.Status(http.StatusOK).JSON(response)
}

func (h Handler) PutTask(c *fiber.Ctx) (err error) {
	// input log
	utils.InputLog(string(c.Context().Path()), c.Request().Header.String(), string(c.Body()))

	// validar token
	if len(c.GetReqHeaders()["Token"]) == 0 {
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "6", Message: "put task error"})
	}
	token := c.GetReqHeaders()["Token"][0]
	id, email, err := utils.ValidToken(token)
	if err != nil {
		bodyError := entity.AppError{Code: "7", Message: "put task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// pegar id task
	idTask, err := c.ParamsInt("id")
	if err != nil {
		if err != nil {
			bodyError := entity.AppError{Code: "9", Message: "put task error"}
			bodyErrorS, _ := json.Marshal(bodyError)
			utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
			return c.Status(http.StatusBadRequest).JSON(bodyError)
		}
	}

	// pegar body
	task := entity.Task{}
	err = c.BodyParser(&task)
	if err != nil {
		bodyError := entity.AppError{Code: "8", Message: "put task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// atualizar task - validar dados usuario
	err = pkg.PutTask(int64(idTask), id, email, task, h.db)
	if err != nil {
		bodyError := entity.AppError{Code: "10", Message: "put task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// response
	utils.ResponseLog("", http.StatusNoContent)
	return c.Status(http.StatusNoContent).SendString("")
}

func (h Handler) DeleteTask(c *fiber.Ctx) (err error) {
	// input log
	utils.InputLog(string(c.Context().Path()), c.Request().Header.String(), string(c.Body()))

	// validar token
	if len(c.GetReqHeaders()["Token"]) == 0 {
		return c.Status(http.StatusBadRequest).JSON(entity.AppError{Code: "6", Message: "delete task error"})
	}
	token := c.GetReqHeaders()["Token"][0]
	id, _, err := utils.ValidToken(token)
	if err != nil {
		bodyError := entity.AppError{Code: "7", Message: "delete task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// pegar id task
	idTask, err := c.ParamsInt("id")
	if err != nil {
		if err != nil {
			bodyError := entity.AppError{Code: "9", Message: "delete task error"}
			bodyErrorS, _ := json.Marshal(bodyError)
			utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
			return c.Status(http.StatusBadRequest).JSON(bodyError)
		}
	}

	// delete task
	err = pkg.DeleteTask(int64(idTask), id, h.db)
	if err != nil {
		bodyError := entity.AppError{Code: "10", Message: "delete task error"}
		bodyErrorS, _ := json.Marshal(bodyError)
		utils.ResponseError(err.Error(), string(bodyErrorS), http.StatusBadRequest)
		return c.Status(http.StatusBadRequest).JSON(bodyError)
	}

	// response
	utils.ResponseLog("", http.StatusNoContent)
	return c.Status(http.StatusNoContent).SendString("")
}
