package ports

import (
	"github.com/edlingao/hexago/internal/calculator/core"
	"github.com/goforj/godump"
)

const (
	ADDITION       = 0
	SUBSTRACTION   = 1
	MULTIPLICATION = 2
	DIVISION       = 3
)

type CalculatorOperations interface {
	Calculate(number1, number2, operation int) int
	Add(number1, number2 int) int
	Substract(number1, number2 int) int
	Multiply(number1, number2 int) int
	Divide(number1, number2 int) int
	GetSymbol(operation int) string
}

type Calculator struct {
	DBService StoringOperations[core.Calculation]
}

func NewCalculator(dbService StoringOperations[core.Calculation]) *Calculator {
	return &Calculator{
		DBService: dbService,
	}
}

func (c Calculator) saveCalculation(result, num1, num2, operation int) (core.Calculation, error) {
	calculation := core.Calculation{
		Result:            result,
		CalculationSymbol: c.GetSymbol(operation),
		Number1:           num1,
		Number2:           num2,
	}

	error := c.DBService.Insert(calculation)

	if error != nil {
		godump.Dump(error)
		return core.Calculation{}, error
	}

	return calculation, nil
}

func (c Calculator) GetSymbol(operation int) string {
	switch operation {
	case ADDITION:
		return "+"
	case SUBSTRACTION:
		return "-"
	case MULTIPLICATION:
		return "*"
	case DIVISION:
		return "/"
	default:
		return ""
	}
}

func (c Calculator) Calculate(num1, num2, operation int) int {
	switch operation {
	case ADDITION:
		return c.Add(num1, num2)
	case SUBSTRACTION:
		return c.Substract(num1, num2)
	case MULTIPLICATION:
		return c.Multiply(num1, num2)
	case DIVISION:
		return c.Divide(num1, num2)
	default:
		return 0
	}
}

func (c Calculator) Add(num1, num2 int) int {
	result := num1 + num2
	c.saveCalculation(result, num1, num2, ADDITION)
	return result
}

func (c Calculator) Substract(num1, num2 int) int {
	result := num1 - num2
	c.saveCalculation(result, num1, num2, SUBSTRACTION)
	return result
}

func (c Calculator) Multiply(num1, num2 int) int {
	result := num1 * num2
	c.saveCalculation(result, num1, num2, MULTIPLICATION)
	return result
}

func (c Calculator) Divide(num1, num2 int) int {
	result := num1 / num2
	c.saveCalculation(result, num1, num2, DIVISION)
	return result
}

func (c Calculator) GetAllCalculations() []core.Calculation {
	return c.DBService.GetAll("calculations")
}
