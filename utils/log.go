package utils

import "github.com/gofiber/fiber/v2/log"

func InputLog(rid string, path string, header string, body string) {
	log.Debugf("[%s] route: %s", rid, path)
	log.Debugf("[%s] request header: %s", rid, header)
	log.Debugf("[%s] request body: %s", rid, body)
}

func ResponseLog(rid string, body string, httpStatus int) {
	if body == "" {
		body = "empty"
	}
	log.Debugf("[%s] response body: %s", rid, body)
	log.Debugf("[%s] response http status: %d", rid, httpStatus)
}

func ResponseError(rid string, err string, body string, httpStatus int) {
	log.Errorf("[%s] response error: %s", rid, err)
	log.Errorf("[%s] response error - body: %s", rid, body)
	log.Errorf("[%s] response error - http status: %d", rid, httpStatus)
}
