package pkg

import (
	"database/sql"
	"errors"

	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/utils"
)

func GetAllTasks(idTask int64, id int64, email string, db *sql.DB) (tasks []entity.Task, err error) {
	// listar todas as tasks do usuário
	tasks, err = utils.SelectTasks(idTask, id, email, db)
	return
}

func GetTask(idTask int64, id int64, email string, db *sql.DB) (task entity.Task, err error) {
	// listar todas as tasks do usuário
	task, err = utils.SelectTask(idTask, id, email, db)
	return
}

func PostTask(id int64, email string, task entity.Task, db *sql.DB) (idTask int64, err error) {
	// validar entrada
	err = fieldTaskValidation(task, "post")
	if err != nil {
		return
	}

	// guardar task - validar dados usuario
	idTask, err = utils.InsertTask(id, email, task, db)
	return
}

func PutTask(idTask int64, id int64, email string, task entity.Task, db *sql.DB) (err error) {
	// validar entrada
	err = fieldTaskValidation(task, "put")
	if err != nil {
		return
	}

	// guardar task - validar dados usuario
	err = utils.UpdateTask(idTask, id, email, task, db)
	return
}

func DeleteTask(idTask int64, id int64, db *sql.DB) (err error) {
	// deletar task
	err = utils.DeleteTask(idTask, id, db)
	return
}

func fieldTaskValidation(task entity.Task, operation string) error {
	if task.Titulo == "" && (operation == "post" || operation == "put") {
		return errors.New("missing field Titulo")
	} else if task.Descricao == "" && (operation == "post" || operation == "put") {
		return errors.New("missing field Descricao")
	} else if task.Status == "" && operation == "put" {
		return errors.New("missing field Status")
	}
	return nil
}
