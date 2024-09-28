package service

import (
	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
)

type RoleService struct {
	roleRepository repository.RoleRepo
}

func NewRoleService(
	roleRepository repository.RoleRepo,
) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (s *RoleService) CreateRole(name string) error {
	role := &entity.Role{
		Name: name,
	}
	err := s.roleRepository.Store(role)
	return err
}

func (s *RoleService) GetAllRoles() []entity.Role {
	return s.roleRepository.ListAll()
}

func (s *RoleService) GetRoleByName(name string) *entity.Role {
	return s.roleRepository.GetByName(name)
}

func (s *RoleService) CountAllRoles() int64 {
	return s.roleRepository.Count()
}
