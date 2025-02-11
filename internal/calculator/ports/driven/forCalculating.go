package driven

type CalculatorOperations interface {
  Calculate(number1, number2, operation int) int
  Add(number1, number2 int) int
  Substract(number1, number2 int) int
  Multiply(number1, number2 int) int
  Divide(number1, number2 int) int
  GetSymbol(operation int) string
}

