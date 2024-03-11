package pkg

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rafaeldajuda/tech-task-golang-api/entity"
	"github.com/rafaeldajuda/tech-task-golang-api/utils"
)

func Login(user entity.User, db *sql.DB) (token string, err error) {
	// validar campos
	err = fieldUserValidation(user, "login")
	if err != nil {
		return
	}

	// validar no banco de dados
	id, exist, err := utils.GetUser(user.Email, user.Senha, db)
	if err != nil {
		return
	}
	if !exist {
		return "", errors.New("user not exist")
	}

	// gerar token
	user.ID = id
	token, err = utils.GenToken(user)
	if err != nil {
		return "", fmt.Errorf("error gen token: %s", err.Error())
	}

	return
}

func Register(user entity.User, db *sql.DB) (err error) {
	err = fieldUserValidation(user, "register")
	if err != nil {
		return
	}

	// validar no banco de dados
	_, exist, err := utils.GetUser(user.Email, user.Senha, db)
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

func fieldUserValidation(user entity.User, operation string) error {
	if user.Nome == "" && operation == "register" {
		return errors.New("missing field Nome")
	} else if user.Email == "" {
		return errors.New("missing field Email")
	} else if user.Senha == "" {
		return errors.New("missing field Senha")
	}
	return nil
}
