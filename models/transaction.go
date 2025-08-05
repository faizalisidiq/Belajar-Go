package models

import (
	"time"
)

type Transaction struct {
	ID        uint       `gorm:"primaryKey"`
	ItemID    uint
	ItemName  string
	Quantity  int
	Total     float64
	Timestamp time.Time `gorm:"autoCreateTime"`
}