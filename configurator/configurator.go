package configurator

import (
	"os"

	"github.com/edlingao/hexago/internal/calculator/adapters"
	"github.com/edlingao/hexago/internal/calculator/core"
	"github.com/labstack/echo/v4"
  _ "github.com/joho/godotenv/autoload"
)

type Configurator struct {
  CalculatorHandler adapters.CalculatorHandler
  CalculatorWebPage adapters.CalculatorWebpage
  Echo *echo.Echo
  v1 *echo.Group
  root *echo.Group
}

func New(
  echo *echo.Echo,
) *Configurator {
  // V1
  api := echo.Group("/api")
  v1 := api.Group("/v1")

  root := echo.Group("")
  
	return &Configurator{
    Echo: echo,
    v1: v1,
    root: root,
  }
}

func (c *Configurator) AddCalculatorAPI() *Configurator {
  dbService := adapters.NewDB[core.Calculation]()
  calcService := core.NewCalculator(dbService)
  calculatorHttpService := c.v1.Group("/calculator")
  
  calculationHandler := adapters.NewCalculatorHandler(
    "/calculator",
    calculatorHttpService,
    calcService,
    dbService,
  )

  c.CalculatorHandler = *calculationHandler

  return c
}

func (c *Configurator) AddCalculatorWeb() *Configurator {
  dbService := adapters.NewDB[core.Calculation]()
  calcService := core.NewCalculator(dbService)
  calculatorWebpageService := adapters.NewCalculatorWebpage(
    "/",
    c.root,
    calcService,
    dbService,
  )

  c.CalculatorWebPage = *calculatorWebpageService
  return c
}

func (c *Configurator) Start() {
  port := os.Getenv("GO_PORT")
  c.Echo.Logger.Fatal(c.Echo.Start(":" + port))
}
