package seeders

import (
	accountTypes "github.com/MelvinNunes/menuz-go/internal/app/constants/account_types"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"github.com/MelvinNunes/menuz-go/internal/domain/service"
	"gorm.io/gorm"
)

func seedRoles(db *gorm.DB) {
	// init repository and service
	roleRepo := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(*roleRepo)
	// end repository and service initialization

	roles := []string{accountTypes.ADMIN, accountTypes.USER}
	totalRoles := roleService.CountAllRoles()
	if totalRoles == 0 {
		for _, roleName := range roles {
			roleService.CreateRole(roleName)
		}
	}
}
