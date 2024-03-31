package queries

import (
	"fmt"
	"reflect"

	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/database"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/dtos"
	"github.com/MelvinNunes/menuz-go/internal/infrastructure/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

var _ repository.UserRepository = &UserRepo{}

func (r *UserRepo) CreateUser(user *entity.User) (*entity.User, error) {
	err := r.db.Clauses(clause.Returning{}).Create(user).Error
	if err != nil {
		return nil, err
	}
	var finalUser entity.User
	r.db.Where("email = ?", user.Email).First(&finalUser)
	return &finalUser, nil
}

func (r *UserRepo) UpdateUser(userID uuid.UUID, userData dtos.UpdateUserDTO) (bool, error) {
	updatable := []string{"device_type", "device_firebase_token", "stripe_customer_id", "username", "identity_number", "password", "activated_at", "identify_status", "category", "customer_registration_status", "merchant_role", "merchant_nationality"}

	updateData := make(map[string]interface{})

	userDataValue := reflect.ValueOf(userData)
	userDataType := reflect.TypeOf(userData)

	for i := 0; i < userDataValue.NumField(); i++ {
		fieldName := userDataType.Field(i).Name
		fieldValue := userDataValue.Field(i).Interface()
		for _, v := range updatable {
			if fieldName == v {
				updateData[util.StringToSnakeCase(fieldName)] = fieldValue
				break
			}
		}
	}

	err := r.db.Where(&entity.User{
		ID: userID,
	}).Updates(updateData).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepo) GetUserByID(ID uuid.UUID) *entity.User {
	var user entity.User

	err := r.db.Where(entity.User{
		ID: ID,
	}).First(&user).Error

	if err != nil {
		return nil
	}

	return &user
}

func (r *UserRepo) GetUserByEmail(email string) *entity.User {
	var user entity.User

	err := r.db.Where(entity.User{
		Email: email,
	}).First(&user).Error

	if err != nil {
		return nil
	}
	return &user
}

func (r *UserRepo) GetUserByPhone(phoneNumberCode string, phoneNumber string) *entity.User {
	var user entity.User

	phone := fmt.Sprintf("%v%v", phoneNumberCode, phoneNumber)

	err := r.db.Where(entity.User{
		PhoneNumber: phone,
	}).First(&user).Error

	if err != nil {
		return nil
	}

	return &user
}

func (r *UserRepo) GetUserByCode(code string) *entity.User {
	var user entity.User
	err := r.db.Where(entity.User{
		Code: code,
	}).First(&user).Error

	if err != nil {
		return nil
	}

	return &user
}

func (r *UserRepo) GetUserByPhoneWithCode(phoneNumber string) *entity.User {
	var user entity.User

	err := r.db.Where(entity.User{
		PhoneNumber: phoneNumber,
	}).First(&user).Error

	if err != nil {
		return nil
	}

	return &user
}

func (r *UserRepo) GetAllUsers(pageNo int, itemsPerPage int) []entity.User {
	var users []entity.User
	r.db.Scopes(database.Paginate(pageNo, itemsPerPage)).Find(&users)
	return users
}

func (r *UserRepo) GetTotalUsers() int64 {
	var total int64
	r.db.Model(&entity.User{}).Count(&total)
	return total
}

func (r *UserRepo) UserExistsByPhone(phoneNumberCode string, phoneNumber string) bool {
	var exists bool
	r.db.Model(&entity.User{}).
		Select("COUNT(*) > 0").
		Where("raw_phone_number = ?", fmt.Sprintf("%v%v", phoneNumberCode, phoneNumber)).
		Find(&exists)
	return exists
}

func (r *UserRepo) UserExistsByApiKey(apiKey string) bool {
	var exists bool
	r.db.Model(&entity.User{}).
		Select("COUNT(*) > 0").
		Where("api_key = ?", apiKey).
		Find(&exists)
	return exists
}

func (r *UserRepo) UserExistsByEmail(email string) bool {
	var exists bool

	r.db.Model(&entity.User{}).
		Select("COUNT(*) > 0").
		Where("email = ?", email).
		Find(&exists)

	return exists
}
