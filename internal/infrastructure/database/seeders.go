package database

import (
	"log"
	"os"

	accountTypes "github.com/MelvinNunes/menuz-go/internal/app/constants/account_types"
	appLanguages "github.com/MelvinNunes/menuz-go/internal/app/constants/app_languages"
	phoneCodes "github.com/MelvinNunes/menuz-go/internal/app/constants/phone_codes"
	"github.com/MelvinNunes/menuz-go/internal/domain/service"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
)

func SeedDatabase() {
	seedRoles()
	seedAdmin()
}

func seedRoles() {
	roles := []string{accountTypes.ADMIN, accountTypes.USER}
	totalRoles := service.RoleService.CountAllRoles()
	if totalRoles == 0 {
		for _, roleName := range roles {
			service.RoleService.CreateRole(roleName)
		}
	}
}

func seedAdmin() {
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

	if !service.UserService.UserExistsByEmail(email) {
		account := dtos.CreateAccount{
			Email:           email,
			Password:        password,
			FirstName:       firstName,
			LastName:        lastName,
			PhoneNumber:     phoneNumber,
			PhoneNumberCode: phoneCodes.MOZAMBIQUE,
			AppLanguage:     appLanguages.EN,
		}

		role := service.RoleService.GetRoleByName(accountTypes.ADMIN)
		if role == nil {
			log.Fatal("Admin role not found in database (infrastructure/database/seeders.go)!")
		}

		service.AccountService.CreateAccount(&account, *role)
	}
}
