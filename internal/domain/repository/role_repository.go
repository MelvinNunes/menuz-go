package repository

import "github.com/MelvinNunes/menuz-go/internal/domain/entity"

type RoleRepository interface {
	Store(role *entity.Role) error
	GetByName(name string) *entity.Role
	ListAll() []entity.Role
	Count() int64
}
