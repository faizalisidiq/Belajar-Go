package models

type Item struct {
	ID          uint 	`gorm:"primaryKey"`
	Name        string 	`gorm:"not null"`
	Price		float64 `gor:"not null"`
	Quantity	int		`gor:"not null"`
}