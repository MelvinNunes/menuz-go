package repository

import "gorm.io/gorm"

type Repositories struct {
	RoleRepo     *RoleRepo
	UserRepo     *UserRepo
	UserRoleRepo *UserRoleRepo
	ProfileRepo  *ProfileRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		RoleRepo:     NewRoleRepository(db),
		UserRepo:     NewUserRepository(db),
		UserRoleRepo: NewUserRoleRepository(db),
		ProfileRepo:  NewProfileRepository(db),
	}
}
