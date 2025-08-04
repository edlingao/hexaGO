package core

type Calculation struct {
	ID                string `json:"id" db:"id"`
	Result            int    `json:"result" db:"result"`
	CalculationSymbol string `json:"symbol" db:"symbol"`
	Number1           int    `json:"number1" db:"num1"`
	Number2           int    `json:"number2" db:"num2"`
	CreatedAt         string `json:"created_at" db:"created_at"`
}

