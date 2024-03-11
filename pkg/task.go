package pkg

import (
	"errors"
	"fmt"

	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/utils"
)

func GetTask(token string) (id int64, email string, err error) {
	id, email, err = utils.ValidToken(token)
	if err != nil {
		err = fmt.Errorf("error get task: %s", err.Error())
	}
	return
}

func PostTask(id int64, email string, task entity.Task) (err error) {
	// validar entrada
	err = fieldTaskValidation(task, "post")
	if err != nil {
		return
	}

	// guardar task - validar dados usuario

	return
}

func fieldTaskValidation(task entity.Task, operation string) error {
	if task.Titulo == "" && operation == "post" {
		return errors.New("missing field Titulo")
	} else if task.Descricao == "" && operation == "post" {
		return errors.New("missing field Descricao")
	}
	return nil
}
