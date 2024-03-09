package entity

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
