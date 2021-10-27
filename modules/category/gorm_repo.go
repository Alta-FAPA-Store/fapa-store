package category

import (
	"time"

	"gorm.io/gorm"
)

type CategoryTable struct {
	Id        int       `gorm:"id"`
	Name      int       `gorm:"name"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
	DeletedAt gorm.DeletedAt
}
