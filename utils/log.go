package utils

import "github.com/gofiber/fiber/v2/log"

func InputLog(path string, header string, body string) {
	log.Debugf("route: %s", path)
	log.Debugf("request header: %s", header)
	log.Debugf("request body: %s", body)
}

func ResponseLog(body string, httpStatus int) {
	if body == "" {
		body = "empty"
	}
	log.Debugf("response body: %s", body)
	log.Debugf("response http status: %d", httpStatus)
}

func ResponseError(err string, body string, httpStatus int) {
	log.Errorf("response error: %s", err)
	log.Errorf("response error - body: %s", body)
	log.Errorf("response error - http status: %d", httpStatus)
}
