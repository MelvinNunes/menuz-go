package seeders

import "gorm.io/gorm"

func SeedDatabase(db *gorm.DB) {
	seedRoles(db)
	seedAdmin(db)
}
