package entity

type Config struct {
	AppName string
	AppHost string
	AppPort string

	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	AppRoutes AppRoutes
}

type AppRoutes struct {
	PostLogin    string
	PostRegister string
	GetTasks     string
	GetTask      string
	PostTask     string
	PutTask      string
	DeleteTask   string
}

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
