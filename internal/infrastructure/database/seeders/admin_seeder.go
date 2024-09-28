package seeders

import (
	"log"
	"os"

	accountTypes "github.com/MelvinNunes/menuz-go/internal/app/constants/account_types"
	appLanguages "github.com/MelvinNunes/menuz-go/internal/app/constants/app_languages"
	phoneCodes "github.com/MelvinNunes/menuz-go/internal/app/constants/phone_codes"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"github.com/MelvinNunes/menuz-go/internal/domain/service"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
	"gorm.io/gorm"
)

func seedAdmin(db *gorm.DB) {
	email := os.Getenv("ADMIN_EMAIL")
	if email == "" {
		log.Fatal("Please set ADMIN_EMAIL environment variable")
	}

	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		log.Fatal("Please set ADMIN_PASSWORD environment variable")
	}

	firstName := os.Getenv("ADMIN_FIRST_NAME")
	if firstName == "" {
		log.Fatal("Please set ADMIN_FIRST_NAME environment variable")
	}

	lastName := os.Getenv("ADMIN_LAST_NAME")
	if lastName == "" {
		log.Fatal("Please set ADMIN_LAST_NAME environment variable")
	}

	phoneNumber := os.Getenv("ADMIN_PHONE_NUMBER")
	if phoneNumber == "" {
		log.Fatal("Please set ADMIN_PHONE_NUMBER environment variable")
	}

	// init repository and service
	roleRepo := repository.NewRoleRepository(db)
	userRepo := repository.NewUserRepository(db)
	userProfileRepo := repository.NewProfileRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)

	roleService := service.NewRoleService(*roleRepo)
	userService := service.NewUserService(*userRepo, *userProfileRepo, *userRoleRepo)
	accountService := service.NewAccountService(*userRepo, *roleRepo, *userProfileRepo, *userRoleRepo)
	// end repository and service initialization

	if !userService.UserExistsByEmail(email) {
		account := dtos.CreateAccount{
			Email:           email,
			Password:        password,
			FirstName:       firstName,
			LastName:        lastName,
			PhoneNumber:     phoneNumber,
			PhoneNumberCode: phoneCodes.MOZAMBIQUE,
			AppLanguage:     appLanguages.EN,
		}

		role := roleService.GetRoleByName(accountTypes.ADMIN)
		if role == nil {
			log.Fatal("Admin role not found in database (infrastructure/database/seeders.go)!")
		}

		accountService.CreateAccount(&account, *role)
	}
}
