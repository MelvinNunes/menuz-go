package service

import "github.com/MelvinNunes/menuz-go/internal/domain/repository"

type Services struct {
	RoleService    *RoleService
	UserService    *UserService
	AccountService *AccountService
	AuthService    *AuthService
}

func InitServices(repos *repository.Repositories) *Services {
	userService := NewUserService(*repos.UserRepo, *repos.ProfileRepo, *repos.UserRoleRepo)

	return &Services{
		RoleService:    NewRoleService(*repos.RoleRepo),
		UserService:    userService,
		AccountService: NewAccountService(*repos.UserRepo, *repos.RoleRepo, *repos.ProfileRepo, *repos.UserRoleRepo),
		AuthService:    NewAuthService(*userService),
	}
}
