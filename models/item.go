package models

type Item struct {
	ID          int `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Price		float64
	Quantity	int
}