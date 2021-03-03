package server

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/huyhvq/betting/pkg/handler"
	"github.com/huyhvq/betting/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"math"
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
		if err := s.server.Start(":8080"); err != nil && err != http.ErrServerClosed {
			s.server.Logger.Fatal("shutting down the server, ", err.Error())
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

func NewServer(wr repository.WagerRepository) Server {
	e := echo.New()
	validate := validator.New()
	validate.RegisterValidation("two-decimal-places", ValidateTwoDecimalPlaces)

	e.Validator = &CustomValidator{validator: validate}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handler.NewHandler(wr)
	e.POST("/wagers", h.CreateWager)
	e.POST("/buy/:id", h.BuyWager)
	e.GET("/wagers", h.ListWager)
	return &svr{server: e}
}

func ValidateTwoDecimalPlaces(fl validator.FieldLevel) bool {
	value := fl.Field().Float()
	valuef := value * float64(math.Pow(10.0, float64(2)))
	extra := valuef - float64(int(valuef))
	return extra == 0
}
