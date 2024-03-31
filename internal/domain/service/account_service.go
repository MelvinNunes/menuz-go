package service

import (
	"fmt"
	"strings"

	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/security"
	"github.com/google/uuid"
)

type accService struct {
	userRepository        repository.UserRepository
	roleRepository        repository.RoleRepository
	userProfileRepository repository.ProfileRepository
	userRoleRepository    repository.UserRoleRepository
}

var AccountService *accService

func NewAccountService(
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
	userProfileRepository repository.ProfileRepository,
	userRoleRepository repository.UserRoleRepository,
) *accService {
	AccountService = &accService{
		userRepository:        userRepository,
		roleRepository:        roleRepository,
		userProfileRepository: userProfileRepository,
		userRoleRepository:    userRoleRepository,
	}
	return &accService{
		userRepository:        userRepository,
		roleRepository:        roleRepository,
		userProfileRepository: userProfileRepository,
		userRoleRepository:    userRoleRepository,
	}
}

func (s *accService) CreateAccount(dto *dtos.CreateAccount, role entity.Role) error {
	password, err := security.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	code := uuid.NewString()
	uuid, _ := uuid.NewUUID()

	rawApiKey, apiKey := security.GenerateApiKeyJWTtoken(uuid.String())

	for {
		if !s.userRepository.UserExistsByApiKey(apiKey) {
			break
		}
		rawApiKey, apiKey = security.GenerateAPIKey()
	}

	userDTO := &entity.User{
		ID:              uuid,
		Code:            strings.Replace(code, "-", "", -1),
		Email:           dto.Email,
		Password:        password,
		PhoneNumberCode: dto.PhoneNumberCode,
		PhoneNumber:     dto.PhoneNumber,
		RawPhoneNumber:  fmt.Sprintf("%v%v", dto.PhoneNumberCode, dto.PhoneNumber),
		AppLanguage:     dto.AppLanguage,
		Active:          true,
	}

	user, err := s.userRepository.CreateUser(userDTO)
	if err != nil {
		return err
	}

	userProfile := &entity.Profile{
		UserID:      user.ID,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Gender:      dto.Gender,
		Avatar:      nil,
		DateOfBirth: dto.DateOfBirth,
	}
	err = s.userProfileRepository.Create(userProfile)
	if err != nil {
		return err
	}

	userRole := &entity.UserRole{
		UserID: user.ID,
		RoleID: role.ID,
	}

	err = s.userRoleRepository.Create(userRole)
	if err != nil {
		return err
	}

	fmt.Print(rawApiKey)

	return nil
}
