package handler

import (
	"fmt"
	"github.com/huyhvq/betting/pkg/model"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"time"
)

type (
	CreateWagerRequest struct {
		TotalWagerValue   int64   `json:"total_wager_value" validate:"required,gt=0"`
		Odds              int64   `json:"odds" validate:"required,gt=0"`
		SellingPercentage int64   `json:"selling_percentage" validate:"required,gte=1,lte=100"`
		SellingPrice      float64 `json:"selling_price" validate:"required,gt=0,two-decimal-places"`
	}

	CreateWagerResponse struct {
		ID                  int64     `json:"id"`
		TotalWagerValue     int64     `json:"total_wager_value"`
		Odds                int64     `json:"odds"`
		SellingPercentage   int64     `json:"selling_percentage"`
		SellingPrice        float64   `json:"selling_price"`
		CurrentSellingPrice float64   `json:"current_selling_price"`
		PercentageSold      int64     `json:"percentage_sold"`
		AmountSold          float64   `json:"amount_sold"`
		PlacedAt            time.Time `json:"placed_at"`
	}

	Error struct {
		Message string `json:"message"`
	}
)

func (h *handler) CreateWager(c echo.Context) error {
	u := new(CreateWagerRequest)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	minSelling := float64(u.TotalWagerValue) * (float64(u.SellingPercentage) / 100)
	minSelling = math.Round(minSelling*100) / 100
	if u.SellingPrice < minSelling {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid selling_price, selling_price must greater than: %.2f", minSelling))
	}

	wm := model.Wager{
		TotalWagerValue:   u.TotalWagerValue,
		Odds:              u.Odds,
		SellingPercentage: u.SellingPercentage,
		SellingPrice:      u.SellingPrice,
	}
	wm, err := h.wr.Create(wm)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resp := CreateWagerResponse{
		ID:                  wm.ID,
		TotalWagerValue:     wm.TotalWagerValue,
		Odds:                wm.Odds,
		SellingPercentage:   wm.SellingPercentage,
		SellingPrice:        wm.SellingPrice,
		CurrentSellingPrice: wm.CurrentSellingPrice,
		PercentageSold:      wm.PercentageSold,
		AmountSold:          wm.AmountSold,
		PlacedAt:            wm.PlacedAt,
	}
	return c.JSON(http.StatusOK, resp)
}
