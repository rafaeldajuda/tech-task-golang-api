package pkg

import (
	"database/sql"
	"errors"

	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/utils"
)

func GetAllTasks(rid string, idTask int64, id int64, email string, db *sql.DB) (tasks []entity.Task, err error) {
	// listar todas as tasks do usuário
	tasks, err = utils.SelectTasks(rid, idTask, id, email, db)
	return
}

func GetTask(rid string, idTask int64, id int64, email string, db *sql.DB) (task entity.Task, err error) {
	// listar todas as tasks do usuário
	task, err = utils.SelectTask(rid, idTask, id, email, db)
	return
}

func PostTask(rid string, id int64, email string, task entity.Task, db *sql.DB) (idTask int64, err error) {
	// validar entrada
	err = fieldTaskValidation(task, "post")
	if err != nil {
		return
	}

	// guardar task - validar dados usuario
	idTask, err = utils.InsertTask(rid, id, email, task, db)
	return
}

func PutTask(rid string, idTask int64, id int64, email string, task entity.Task, db *sql.DB) (err error) {
	// validar entrada
	err = fieldTaskValidation(task, "put")
	if err != nil {
		return
	}

	// guardar task - validar dados usuario
	err = utils.UpdateTask(rid, idTask, id, email, task, db)
	return
}

func DeleteTask(rid string, idTask int64, id int64, db *sql.DB) (err error) {
	// deletar task
	err = utils.DeleteTask(rid, idTask, id, db)
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
