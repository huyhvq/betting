package server

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server interface {
	Start()
}

type svr struct {
	server *echo.Echo
}

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if cv.validator.Struct(i) != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, cv.validator.Struct(i).Error())
	}
	return nil
}

func (s *svr) Start() {
	go func() {
		if err := s.server.Start(":1323"); err != nil && err != http.ErrServerClosed {
			s.server.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.server.Logger.Fatal(err)
	}
}

func NewServer() Server {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/wagers", createWager)
	return &svr{server: e}
}

