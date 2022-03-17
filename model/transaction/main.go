package transaction

import "time"

//Transaction usada para cadastrar e mostrar uma transação
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

//Transactions lista de transações
type Transactions []Transaction
