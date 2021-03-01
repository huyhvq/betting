package handler

import (
	"github.com/huyhvq/betting/pkg/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	CreateWagerRequest struct {
		TotalWagerValue   int64   `json:"total_wager_value" validate:"required,gt=0"`
		Odds              int64   `json:"odds" validate:"required,gt=0"`
		SellingPercentage int64   `json:"selling_percentage" validate:"required,gte=1,lte=100"`
		SellingPrice      float64 `json:"selling_price" validate:"required,gt=0"`
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

	wm := model.Wager{
		TotalWagerValue:   u.TotalWagerValue,
		Odds:              u.Odds,
		SellingPercentage: u.SellingPercentage,
		SellingPrice:      u.SellingPrice,
	}
	wm, err := h.wr.Create(wm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, wm)
}
