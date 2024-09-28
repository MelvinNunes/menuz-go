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

type AccountService struct {
	userRepository        repository.UserRepo
	roleRepository        repository.RoleRepo
	userProfileRepository repository.ProfileRepo
	userRoleRepository    repository.UserRoleRepo
}

func NewAccountService(
	userRepository repository.UserRepo,
	roleRepository repository.RoleRepo,
	userProfileRepository repository.ProfileRepo,
	userRoleRepository repository.UserRoleRepo,
) *AccountService {
	return &AccountService{
		userRepository:        userRepository,
		roleRepository:        roleRepository,
		userProfileRepository: userProfileRepository,
		userRoleRepository:    userRoleRepository,
	}
}

func (s *AccountService) CreateAccount(dto *dtos.CreateAccount, role entity.Role) error {
	password, err := security.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	code := uuid.NewString()
	uuid, _ := uuid.NewUUID()

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
	return nil
}
