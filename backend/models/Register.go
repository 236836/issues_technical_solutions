package models

import (
	"gorm.io/gorm"
)

type Devo struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required,min=8"`
	Role     string `gorm:"default:user" json:"role"`
}
