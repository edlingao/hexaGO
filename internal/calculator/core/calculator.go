package core

import (
	"log"

	"github.com/edlingao/hexago/internal/calculator/ports/driven"
)

type Calculator struct {
  DBService driven.StoringOperations[Calculation]
}

type Calculation struct {
  ID string `json:"id" db:"id"`
  Result int `json:"result" db:"result"`
  CalculationSymbol string `json:"symbol" db:"symbol"`
  Number1 int `json:"number1" db:"num1"`
  Number2 int `json:"number2" db:"num2"`
  CreatedAt string `json:"created_at" db:"created_at"`
}

const (
  ADDITION = 0
  SUBSTRACTION = 1
  MULTIPLICATION = 2
  DIVISION = 3
)

func NewCalculator(dbService driven.StoringOperations[Calculation]) *Calculator {
  return &Calculator{
    DBService: dbService,
  }
}

func (c Calculator) saveCalculation(result, num1, num2, operation int) ( Calculation, error )  {
  calculation := Calculation{
    Result: result,
    CalculationSymbol: c.GetSymbol(operation),
    Number1: num1,
    Number2: num2,
  }

  error := c.DBService.Insert(calculation, `
    INSERT INTO calculations (result, symbol, num1, num2)
    VALUES (:result, :symbol, :num1, :num2)
  `)

  if error != nil {
    log.Fatal(error)
    return Calculation{}, error
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

