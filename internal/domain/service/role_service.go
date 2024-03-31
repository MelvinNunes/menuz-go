package service

import (
	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
)

type rService struct {
	roleRepository repository.RoleRepository
}

var RoleService *rService

func NewRoleService(
	roleRepository repository.RoleRepository,
) *rService {
	RoleService = &rService{
		roleRepository: roleRepository,
	}
	return &rService{
		roleRepository: roleRepository,
	}
}

func (s *rService) CreateRole(name string) error {
	role := &entity.Role{
		Name: name,
	}
	err := s.roleRepository.Store(role)
	return err
}

func (s *rService) GetAllRoles() []entity.Role {
	return s.roleRepository.ListAll()
}

func (s *rService) GetRoleByName(name string) *entity.Role {
	return s.roleRepository.GetByName(name)
}

func (s *rService) CountAllRoles() int64 {
	return s.roleRepository.Count()
}
