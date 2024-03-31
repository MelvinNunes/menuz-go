package repository

import (
	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserByID(Id uuid.UUID) *entity.User
	GetUserByEmail(email string) *entity.User
	GetUserByCode(code string) *entity.User
	GetUserByPhone(phoneNumberCode string, phoneNumber string) *entity.User
	GetUserByPhoneWithCode(phoneNumber string) *entity.User
	GetAllUsers(pageNo int, itemsPerPage int) []entity.User
	GetTotalUsers() int64
	UserExistsByPhone(phoneNumberCode string, phoneNumber string) bool
	UserExistsByApiKey(apiKey string) bool
	UserExistsByEmail(username string) bool
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(userID uuid.UUID, userData dtos.UpdateUserDTO) (bool, error)
}
