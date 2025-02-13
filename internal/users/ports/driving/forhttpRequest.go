package driving

import "github.com/labstack/echo/v4"

type UserViews interface {
  Home(c echo.Context) error
  Register(c echo.Context) error
  Login(c echo.Context) error
  Settings(c echo.Context) error
}

type Response[T any] struct {
  Status int `json:"status"`
  Message string `json:"message"`
  Data T `json:"data"`
}
