package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	CreateWagerRequest struct {
		TotalWagerValue   int64 `json:"total_wager_value" validate:"required,gt=0"`
		Odds              int64 `json:"odds" validate:"required,gt=0"`
		SellingPercentage int64 `json:"selling_percentage" validate:"required,gte=1,lte=100"`
		SellingPrice      int64 `json:"selling_price" validate:"required,gt=0"`
	}
)

func createWager(c echo.Context) error {
	u := new(CreateWagerRequest)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}
