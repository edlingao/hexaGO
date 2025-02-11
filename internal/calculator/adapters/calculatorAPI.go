package adapters

import (
	"strconv"

	"github.com/edlingao/hexago/internal/calculator/core"
	"github.com/edlingao/hexago/internal/calculator/ports/driven"
	"github.com/edlingao/hexago/internal/calculator/ports/driving"
	"github.com/labstack/echo/v4"
)

type CalculatorHandler struct {
	URL               string
	Group             *echo.Group
	CalculatorService driven.CalculatorOperations
	DBService         driven.StoringOperations[core.Calculation]
}

func NewCalculatorHandler(
  url string,
  httpService *echo.Group,
  calcService driven.CalculatorOperations,
  dbService driven.StoringOperations[core.Calculation],
) *CalculatorHandler {

	calculatorHandler := &CalculatorHandler{
		URL:   url,
		Group: httpService,
    CalculatorService: calcService,
    DBService: dbService,
	}
  
	calculatorHandler.Group.GET("/history", calculatorHandler.GetAllCalculations)
	calculatorHandler.Group.GET("/:id", calculatorHandler.GetCalculation)
	calculatorHandler.Group.DELETE("/delete/:id", calculatorHandler.DeleteCalculation)
	calculatorHandler.Group.POST("/calculate", calculatorHandler.Calculate)

	return calculatorHandler
}

func (ch *CalculatorHandler) GetAllCalculations(c echo.Context) error {
  calculations := ch.DBService.GetAll("calculations")

  return c.JSON(200, driving.Response[[]core.Calculation]{
    Status: 200,
    Message: "Success",
    Data: calculations,
  })
}

func (ch *CalculatorHandler) GetCalculation(c echo.Context) error {
  id := c.Param("id")

  calculation, err := ch.DBService.Get(id, "calculations")

  if err != nil {
    return c.JSON(400, driving.Response[interface{}]{
      Status: 400,
      Message: err.Error(),
      Data: nil,
    })
  }

  return c.JSON(200, driving.Response[core.Calculation]{
    Status: 200,
    Message: "Success",
    Data: calculation,
  })
}

func (ch *CalculatorHandler) DeleteCalculation(c echo.Context) error {
  id := c.Param("id")

  err := ch.DBService.Delete(id, "calculations")

  if err != nil {
    return c.JSON(400, driving.Response[interface{}]{
      Status: 400,
      Message: err.Error(),
      Data: nil,
    })
  }

  return c.JSON(200, driving.Response[interface{}]{
    Status: 200,
    Message: "Success",
    Data: nil,
  })
}

func (ch *CalculatorHandler) Calculate(c echo.Context) error {
  num1, err := strconv.Atoi( c.FormValue("number1") )
  num2, err := strconv.Atoi( c.FormValue("number2") )
  operator, err := strconv.Atoi( c.FormValue("operator") )

  if operator < 0 || operator > 3 {
    return c.JSON(400, driving.Response[interface{}]{
      Status: 400,
      Message: "Invalid operator",
      Data: nil,
    })
  }

  if err != nil {
    return c.JSON(400, driving.Response[interface{}]{
      Status: 400,
      Message: err.Error(),
      Data: nil,
    })
  }

  result := ch.CalculatorService.Calculate(num1, num2, operator)
  
  return c.JSON(200, driving.Response[core.Calculation]{
    Status: 200,
    Message: "Success",
    Data: core.Calculation{
      Result: result,
      CalculationSymbol: ch.CalculatorService.GetSymbol(operator),
      Number1: num1,
      Number2: num2,
    },
  })
}
