package http

import (
	"net/http"

	"github.com/adrianamaiaterosendo/projeto-dio-planejamento-financeiro.git/handler/http/actuator"
	"github.com/adrianamaiaterosendo/projeto-dio-planejamento-financeiro.git/handler/transaction"
)

//Init é onde eu tenho as chamadas http
func Init() {

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transsaction/create", transaction.CreatedTransactions)
	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
