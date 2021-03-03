package model

import (
	"gorm.io/gorm"
)

type Wager struct {
	gorm.Model
	TotalWagerValue     int64
	Odds                int64
	SellingPercentage   int64
	SellingPrice        float64
	CurrentSellingPrice float64
	PercentageSold      int64
	AmountSold          float64
}
