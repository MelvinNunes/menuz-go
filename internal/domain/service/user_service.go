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

type UserService struct {
	userRepository     repository.UserRepo
	profileRepository  repository.ProfileRepo
	userRoleRepository repository.UserRoleRepo
}

func NewUserService(
	userRepository repository.UserRepo,
	profileRepository repository.ProfileRepo,
	userRoleRepository repository.UserRoleRepo,
) *UserService {
	return &UserService{
		userRepository:     userRepository,
		profileRepository:  profileRepository,
		userRoleRepository: userRoleRepository,
	}
}

func (s *UserService) GetUserByID(ID uuid.UUID) *entity.User {
	return s.userRepository.GetUserByID(ID)
}

func (s *UserService) GetUserByEmail(email string) *entity.User {
	return s.userRepository.GetUserByEmail(email)
}

func (s *UserService) GetUserByCode(email string) *entity.User {
	return s.userRepository.GetUserByCode(email)
}

func (s *UserService) UserExistsByEmail(email string) bool {
	return s.userRepository.UserExistsByEmail(email)
}

func (s *UserService) UserExistsByPhoneNumber(phoneCode, phoneNumber string) bool {
	return s.userRepository.UserExistsByPhone(phoneCode, phoneNumber)
}

func (s *UserService) GetUserVMByID(ID uuid.UUID) (*dtos.UserVM, error) {
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

func (s *UserService) GetOnlineUser(c *fiber.Ctx) (*dtos.UserVM, error) {
	userID := security.GetOnlineUserID(c)
	return s.GetUserVMByID(uuid.MustParse(userID))
}
