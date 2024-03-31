package service

import (
	"errors"
	"fmt"

	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/security"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type uService struct {
	userRepository     repository.UserRepository
	profileRepository  repository.ProfileRepository
	userRoleRepository repository.UserRoleRepository
}

var UserService *uService

func NewUserService(
	userRepository repository.UserRepository,
	profileRepository repository.ProfileRepository,
	userRoleRepository repository.UserRoleRepository,
) *uService {
	UserService = &uService{
		userRepository:     userRepository,
		profileRepository:  profileRepository,
		userRoleRepository: userRoleRepository,
	}
	return &uService{
		userRepository:     userRepository,
		profileRepository:  profileRepository,
		userRoleRepository: userRoleRepository,
	}
}

func (s *uService) GetUserByID(ID uuid.UUID) *entity.User {
	return s.userRepository.GetUserByID(ID)
}

func (s *uService) GetUserByEmail(email string) *entity.User {
	return s.userRepository.GetUserByEmail(email)
}

func (s *uService) GetUserByCode(email string) *entity.User {
	return s.userRepository.GetUserByCode(email)
}

func (s *uService) UserExistsByEmail(email string) bool {
	return s.userRepository.UserExistsByEmail(email)
}

func (s *uService) UserExistsByPhoneNumber(phoneCode, phoneNumber string) bool {
	return s.userRepository.UserExistsByPhone(phoneCode, phoneNumber)
}

func (s *uService) GetUserVMByID(ID uuid.UUID) (*dtos.UserVM, error) {
	user := s.userRepository.GetUserByID(ID)
	if user == nil {
		return nil, errors.New("user was not found")
	}

	var roles []string
	userRoles := s.userRoleRepository.GetAllFromUserID(ID)

	for _, userRole := range userRoles {
		roles = append(roles, userRole.Role.Name)
	}

	userVM := dtos.UserVM{
		ID:           user.ID,
		Code:         user.Code,
		Email:        user.Email,
		PhoneNumber:  fmt.Sprintf("%v%v", user.PhoneNumberCode, user.PhoneNumber),
		AccountTypes: roles,
		CreatedAt:    user.CreatedAt.String(),
	}

	return &userVM, nil
}

func (s *uService) GetOnlineUser(c *fiber.Ctx) (*dtos.UserVM, error) {
	userID := security.GetOnlineUserID(c)
	return s.GetUserVMByID(uuid.MustParse(userID))
}
