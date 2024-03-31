package queries

import (
	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleRepo struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) *UserRoleRepo {
	return &UserRoleRepo{db}
}

var _ repository.UserRoleRepository = &UserRoleRepo{}

func (r *UserRoleRepo) Create(userRole *entity.UserRole) error {
	return r.db.Create(userRole).Error
}

func (r *UserRoleRepo) GetAllFromUserID(userID uuid.UUID) []entity.UserRole {
	var userRoles []entity.UserRole
	r.db.Preload("Role").Where("user_id = ?", userID).Find(&userRoles)
	return userRoles
}
