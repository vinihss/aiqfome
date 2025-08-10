package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}
