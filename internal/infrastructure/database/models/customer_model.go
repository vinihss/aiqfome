package models

import (
	"gorm.io/gorm"
	"strings"
)

type Customer struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"uniqueIndex"`
}

func (c *Customer) BeforeSave(tx *gorm.DB) (err error) {
	c.Email = strings.TrimSpace(strings.ToLower(c.Email))
	return nil
}
