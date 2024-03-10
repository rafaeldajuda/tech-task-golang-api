package pkg

import (
	"database/sql"
	"errors"

	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/utils"
)

func Login(user entity.User, db *sql.DB) (token string, err error) {
	// validar campos
	err = fieldValidation(user, "login")
	if err != nil {
		return
	}

	// validar no banco de dados
	exist, err := utils.GetUser(user.Email, user.Senha, db)
	if err != nil {
		return
	}
	if !exist {
		return "", errors.New("user not exist")
	}

	// retornar token
	return "token123", nil
}

func Register(user entity.User, db *sql.DB) (err error) {
	err = fieldValidation(user, "register")
	if err != nil {
		return
	}

	// validar no banco de dados
	exist, err := utils.GetUser(user.Email, user.Senha, db)
	if err != nil {
		return
	}
	if exist {
		return errors.New("user exist")
	}

	// guardar usuário
	err = utils.InsertUser(user.Nome, user.Email, user.Senha, db)
	if err != nil {
		return
	}

	return
}

func fieldValidation(user entity.User, operation string) error {
	if user.Nome == "" && operation == "register" {
		return errors.New("missing field Nome")
	} else if user.Email == "" {
		return errors.New("missing field Email")
	} else if user.Senha == "" {
		return errors.New("missing field Senha")
	}
	return nil
}
