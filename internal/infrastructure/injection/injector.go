package injection

import (
	"github.com/MelvinNunes/menuz-go/internal/domain/service"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/queries"
	"gorm.io/gorm"
)

func InjectReposAndService(db *gorm.DB) {
	// init repository and service
	roleRepo := queries.NewRoleRepository(db)
	userRepo := queries.NewUserRepository(db)
	userProfileRepo := queries.NewProfileRepository(db)
	userRoleRepo := queries.NewUserRoleRepository(db)

	service.NewRoleService(roleRepo)
	service.NewUserService(userRepo, userProfileRepo, userRoleRepo)
	service.NewAccountService(userRepo, roleRepo, userProfileRepo, userRoleRepo)
	// end repository and service initialization
}
