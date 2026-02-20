package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"size:255;unique;not null"`
	Email     string `gorm:"size:255;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
