package models

// Expense representa um gasto público
type Expense struct {
	Codigo    int     `json:"codigo"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
	Data      string  `json:"data"`
}

// APIExpensesResponse representa a resposta da API
type APIExpensesResponse struct {
	Response []Expense `json:"response"`
}
