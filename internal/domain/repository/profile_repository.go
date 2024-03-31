package repository

import "github.com/MelvinNunes/menuz-go/internal/domain/entity"

type ProfileRepository interface {
	Create(userProfile *entity.Profile) error
	Update(profileID uint64, userProfile *entity.Profile) error
}
