package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rafaeldajuda/tech-task-golang-api/entity"
)

const SECRET_KEY = "secret123"

func GenToken(user entity.User) (token string, err error) {
	claims := jwt.MapClaims{
		"Nome": user.Email,
		"exp":  time.Now().Add(time.Minute * 5).Local().Unix(),
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tk.SignedString([]byte(SECRET_KEY))
	return
}

func ValidToken(token string) (err error) {
	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return
	}

	if claims, ok := tk.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["Nome"], claims["exp"])
		exp := int64(claims["exp"].(float64))
		if time.Now().Local().Unix() > exp {
			return errors.New("invalid exp token")
		}
	}
	return
}
