package middleware

import "github.com/MelvinNunes/menuz-go/internal/domain/service"

type Middlewares struct {
}

func InitMiddlewares(services *service.Services) *Middlewares {
	return &Middlewares{}
}
