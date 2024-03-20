package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/handler"
	"github.com/rafaeldajuda/tech-task-golang-api/utils"
)

var config entity.Config

func init() {
	godotenv.Load()

	// load envs
	config = entity.Config{
		AppName:    os.Getenv("APP_NAME"),
		AppHost:    os.Getenv("APP_HOST"),
		AppPort:    os.Getenv("APP_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		AppRoutes: entity.AppRoutes{
			PostLogin:    os.Getenv("POST_LOGIN"),
			PostRegister: os.Getenv("POST_REGISTER"),
			GetTasks:     os.Getenv("GET_TASKS"),
			GetTask:      os.Getenv("GET_TASK"),
			PostTask:     os.Getenv("POST_TASK"),
			PutTask:      os.Getenv("PUT_TASK"),
			DeleteTask:   os.Getenv("DELETE_TASK"),
		},
	}
}

func main() {

	log.Debugf("** STARTING APP %s **", config.AppName)

	// database
	db := utils.Connection(config)
	mapStatus := utils.ConfigTables(db)

	// app
	handler := handler.NewHandler(db, mapStatus)
	app := fiber.New()

	app.Post(config.AppRoutes.PostLogin, handler.PostLogin)
	app.Post(config.AppRoutes.PostRegister, handler.PostRegister)
	app.Get(config.AppRoutes.GetTasks, handler.GetTasks)
	app.Get(config.AppRoutes.GetTask, handler.GetTask)
	app.Post(config.AppRoutes.PostTask, handler.PostTask)
	app.Put(config.AppRoutes.PutTask, handler.PutTask)
	app.Delete(config.AppRoutes.DeleteTask, handler.DeleteTask)

	log.Debug("** STARTING SERVER **")
	host := fmt.Sprintf("%s:%s", config.AppHost, config.AppPort)
	if err := app.Listen(host); err != nil {
		log.Fatalf("error starting server, %s", err.Error())
	}

}
