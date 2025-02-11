package driving

import (
	"github.com/labstack/echo/v4"
)

type HTTPHandlerCalculator interface {
  GetAllCalculations(c echo.Context) error
  GetCalculation(c echo.Context) error
  SaveCalculation(c echo.Context) error
  DeleteCalculation(c echo.Context) error
  Calculate(c echo.Context) error
}

type CalculatorWebViews interface {
  Home(c echo.Context) error
}

type Response[T any] struct {
  Status int `json:"status"`
  Message string `json:"message"`
  Data T `json:"data"`
}

