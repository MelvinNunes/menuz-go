package queries

import (
	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"gorm.io/gorm"
)

type RoleRepo struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepo {
	return &RoleRepo{db}
}

var _ repository.RoleRepository = &RoleRepo{}

func (r *RoleRepo) Store(role *entity.Role) error {
	err := r.db.Create(role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleRepo) GetByName(name string) *entity.Role {
	var role entity.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	if err != nil {
		return nil
	}
	return &role
}

func (r *RoleRepo) ListAll() []entity.Role {
	var roles []entity.Role
	r.db.Find(&roles)
	return roles
}

func (r *RoleRepo) Count() int64 {
	var total int64
	r.db.Model(&entity.Role{}).Count(&total)
	return total
}
