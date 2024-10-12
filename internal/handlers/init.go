package handlers

import "github.com/MelvinNunes/menuz-go/internal/domain/service"

type Handlers struct {
	AccountHandler *AccountHandler
	AuthHandler    *AuthHandler
}

func InitHandlers(services *service.Services) *Handlers {
	return &Handlers{
		AccountHandler: NewAccountHandler(*services.UserService, *services.RoleService, *services.AccountService),
		AuthHandler:    NewAuthHandler(*services.AuthService),
	}
}
