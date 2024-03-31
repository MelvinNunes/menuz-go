package queries

import (
	"github.com/MelvinNunes/menuz-go/internal/domain/entity"
	"github.com/MelvinNunes/menuz-go/internal/domain/repository"
	"gorm.io/gorm"
)

type ProfileRepo struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepo {
	return &ProfileRepo{db}
}

var _ repository.ProfileRepository = &ProfileRepo{}

func (r *ProfileRepo) Create(userProfile *entity.Profile) error {
	return r.db.Create(userProfile).Error
}

func (r *ProfileRepo) Update(profileID uint64, userProfile *entity.Profile) error {
	return nil
}
