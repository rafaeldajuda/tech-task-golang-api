package pkg

import (
	"fmt"

	"github.com/rafaeldajuda/tech-task-golang-api/utils"
)

func GetTask(token string) (err error) {
	err = utils.ValidToken(token)
	if err != nil {
		return fmt.Errorf("error get task: %s", err.Error())
	}
	return
}
