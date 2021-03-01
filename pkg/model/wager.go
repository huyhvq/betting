package model

import "time"

type Wager struct {
	ID                  int64
	TotalWagerValue     int64
	Odds                int64
	SellingPercentage   int64
	SellingPrice        float64
	CurrentSellingPrice float64
	PercentageSold      int64
	AmountSold          float64
	PlacedAt            time.Time
}
