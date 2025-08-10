package models

type Favorite struct {
	ID         uint   `gorm:"primaryKey"`
	CustomerID uint   `gorm:"not null;index;uniqueIndex:uniq_customer_product"`
	ProductID  uint   `gorm:"not null;uniqueIndex:uniq_customer_product"`
	Product    string `gorm:"not null"`
}
