package pkg

import (
	"errors"

	"github.com/rafaeldajuda/tech-task-golang-api/entity"
)

func Login(user entity.User) (token string, err error) {
	// validar campos
	// checar se usuário existe no banco
	// retornar resposta
	return
}

func Register(user entity.User) (err error) {
	err = registerValidation(user)
	if err != nil {
		return
	}
	// err = checkUser(user)
	// if err != nil {
	// 	return
	// }
	// validar no banco de dados
	// guardar usuário
	// retornar resposta
	return
}

func registerValidation(user entity.User) error {
	if user.Nome == "" {
		return errors.New("missing field Nome")
	} else if user.Email == "" {
		return errors.New("missing field Email")
	} else if user.Senha == "" {
		return errors.New("missing field Senha")
	}
	return nil
}
