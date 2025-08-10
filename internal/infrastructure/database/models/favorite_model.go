package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	CustomerID string
	ProductID  string
	Product    string
}
