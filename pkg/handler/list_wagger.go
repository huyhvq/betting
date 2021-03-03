package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *handler) ListWager(c echo.Context) error {
	var page, limit int = 1, 10
	var err error

	if p := c.QueryParam("page"); p != "" {
		fmt.Println(p)
		page, err = strconv.Atoi(p)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid page parameter")
		}
	}
	if p := c.QueryParam("limit"); p != "" {
		limit, err = strconv.Atoi(p)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid limit parameter")
		}
	}

	offset := (page - 1) * limit

	ws, err := h.wr.List(offset, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := make([]CreateWagerResponse, 0, len(ws))
	for _, w := range ws {
		resp = append(resp, CreateWagerResponse{
			ID:                  w.ID,
			TotalWagerValue:     w.TotalWagerValue,
			Odds:                w.Odds,
			SellingPercentage:   w.SellingPercentage,
			SellingPrice:        w.SellingPrice,
			CurrentSellingPrice: w.CurrentSellingPrice,
			PercentageSold:      w.PercentageSold,
			AmountSold:          w.AmountSold,
			PlacedAt:            w.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, resp)
}
