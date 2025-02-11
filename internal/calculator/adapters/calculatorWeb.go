package adapters

import (
	"strconv"

	"github.com/edlingao/hexago/internal/calculator/core"
	"github.com/edlingao/hexago/internal/calculator/ports/driven"
	"github.com/edlingao/hexago/lib"
	"github.com/edlingao/hexago/web/views"
	"github.com/labstack/echo/v4"
)

type CalculatorWebpage struct {
	URL         string
	http        *echo.Group
	calcService driven.CalculatorOperations
	dbService   driven.StoringOperations[core.Calculation]
}

func NewCalculatorWebpage(
	url string,
	httpService *echo.Group,
	calcService driven.CalculatorOperations,
	dbService driven.StoringOperations[core.Calculation],
) *CalculatorWebpage {

	calculatorWebPageService := &CalculatorWebpage{
		URL:         url,
		http:        httpService,
		calcService: calcService,
		dbService:   dbService,
	}

	calculatorWebPageService.http.GET("/", calculatorWebPageService.Home)
	calculatorWebPageService.http.POST("/calculate", calculatorWebPageService.Calculate)
  calculatorWebPageService.http.GET("/history", calculatorWebPageService.History)

	return calculatorWebPageService
}

func (cw *CalculatorWebpage) Home(c echo.Context) error {
	history := cw.dbService.GetAll("calculations")

	return lib.Render(
		c,
		views.Index(views.IndexVM{
			Result:  views.IndexResult{},
			History: history,
		}),
		200,
	)
}

func (cw *CalculatorWebpage) Calculate(c echo.Context) error {
	num1, err := strconv.Atoi(c.FormValue("num1"))
	num2, err := strconv.Atoi(c.FormValue("num2"))
	operation, err := strconv.Atoi(c.FormValue("operation"))

	if err != nil {
		return lib.Render(
			c,
			views.Index(views.IndexVM{
				Error:   err,
				Result:  views.IndexResult{},
				History: cw.dbService.GetAll("calculations"),
			}),
			400,
		)
	}
	c.Response().Header().Set("HX-Trigger", "calc:history")
	result := cw.calcService.Calculate(num1, num2, operation)

	return lib.Render(
		c,
		views.Index(views.IndexVM{
			Result: views.IndexResult{
				Result: result,
			},
			History: cw.dbService.GetAll("calculations"),
		}),
		200,
	)
}

func (cw *CalculatorWebpage) History(c echo.Context) error {
	history := cw.dbService.GetAll("calculations")

	return lib.Render(
		c,
		views.Index(views.IndexVM{
			Result:  views.IndexResult{},
			History: history,
		}),
		200,
	)
}
