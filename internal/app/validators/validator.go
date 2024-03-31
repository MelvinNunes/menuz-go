package validators

import (
	"errors"

	accountTypes "github.com/MelvinNunes/menuz-go/internal/app/constants/account_types"
	appLanguages "github.com/MelvinNunes/menuz-go/internal/app/constants/app_languages"
	"github.com/MelvinNunes/menuz-go/internal/app/constants/currencies"
	phoneCodes "github.com/MelvinNunes/menuz-go/internal/app/constants/phone_codes"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/util"
	"github.com/go-playground/validator/v10"
)

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// This is the validator instance
// use a single instance of Validate, it caches struct info
var Validate *validator.Validate

func EnableValidation() {
	Validate = validator.New(validator.WithRequiredStructEnabled())

	Validate.RegisterValidation("phone_code_validator", isValidPhoneCode)
	Validate.RegisterValidation("lang_validator", isValidLanguage)
	Validate.RegisterValidation("role_validator", isValidRole)
	Validate.RegisterValidation("currency_validator", isValidCurrency)

}

func ValidateData(data interface{}) (*map[string]string, error) {
	errArray := make(map[string]string)
	errs := Validate.Struct(data)

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			errArray[util.StringToSnakeCase(err.Field())] = err.Tag()
		}
		return &errArray, errors.New("validation error")
	}
	return nil, nil
}

func isValidPhoneCode(fl validator.FieldLevel) bool {
	return fl.Field().String() == phoneCodes.MOZAMBIQUE
}

func isValidLanguage(fl validator.FieldLevel) bool {
	return fl.Field().String() == appLanguages.EN || fl.Field().String() == appLanguages.PT
}

func isValidRole(fl validator.FieldLevel) bool {
	return fl.Field().String() == accountTypes.ADMIN || fl.Field().String() == accountTypes.USER
}

func isValidCurrency(fl validator.FieldLevel) bool {
	return fl.Field().String() == currencies.MZN
}
