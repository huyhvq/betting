package handler

import (
	"github.com/huyhvq/betting/pkg/repository"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	CreateWager(c echo.Context) error
}

type handler struct {
	wr repository.WagerRepository
}

func NewHandler(wr repository.WagerRepository) Handler {
	return &handler{wr: wr}
}
