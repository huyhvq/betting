package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	WagerID     uint
	BuyingPrice float64
}
