package main

import (
	"embed"
	"net/http"

	"github.com/edlingao/hexago/configurator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed all:static
var static embed.FS

func main() {
	echo := echo.New()
	echo.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "/",
		Browse:     false,
		HTML5:      false,
		Filesystem: http.FS(static),
	}))
	config := configurator.New(
		echo,
	)

	config.AddCalculatorAPI()
  config.AddCalculatorWeb()
  config.AddUserAPI()
  config.AddUserWeb()
  config.Start()
}
