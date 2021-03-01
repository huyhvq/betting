package handler

import "github.com/labstack/echo/v4"

type Handler interface {
	CreateWager(c echo.Context) error
}

type handler struct {
}

func NewHandler(db *) Handler {

}
