package repository

import (
	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRoleRepository interface {
	Create(userRole *entity.UserRole) error
	GetAllFromUserID(userID uuid.UUID) []entity.UserRole
}
