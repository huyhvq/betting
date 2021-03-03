package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type (
	BuyWagerRequest struct {
		BuyingPrice float64 `json:"buying_price" validate:"required,gt=0,two-decimal-places"`
	}

	BuyWagerResponse struct {
		ID          uint      `json:"id"`
		WagerID     uint      `json:"wager_id"`
		BuyingPrice float64   `json:"buying_price"`
		BoughtAt    time.Time `json:"bought_at"`
	}
)

func (h *handler) BuyWager(c echo.Context) error {
	req := new(BuyWagerRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 0, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid wager id")
	}
	w, err := h.wr.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if req.BuyingPrice > w.CurrentSellingPrice {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid buying_price")
	}
	t, err := h.wr.Buy(uint(id), req.BuyingPrice)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resp := &BuyWagerResponse{
		ID:          t.ID,
		WagerID:     t.WagerID,
		BuyingPrice: t.BuyingPrice,
		BoughtAt:    t.CreatedAt,
	}
	return c.JSON(http.StatusOK, resp)
}
