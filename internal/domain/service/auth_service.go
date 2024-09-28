package service

import (
	"errors"

	errorchecker "github.com/MelvinNunes/menuz-go/internal/infrastructure/error"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/security"
)

type LoginResult struct {
	Token string `json:"token"`
}

type AuthService struct {
	userService UserService
}

func NewAuthService(userService UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

func (s *AuthService) Login(email, password string) (*LoginResult, error) {
	user := s.userService.GetUserByEmail(email)
	if user == nil {
		return nil, errors.New(errorchecker.UserNotFound)
	}
	if !security.IsPasswordEqualToHash(password, user.Password) {
		return nil, errors.New(errorchecker.InvalidPassword)
	}
	token, err := security.CreateJWTtoken(user.ID.String())

	if err != nil {
		return nil, errors.New(errorchecker.InternalErrorCreatingJWT)
	}

	res := LoginResult{
		Token: *token,
	}

	return &res, nil
}
