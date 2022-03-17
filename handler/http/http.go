package http

import (
	"net/http"

	"github.com/adrianamaiaterosendo/projeto-dio-planejamento-financeiro.git/handler/http/actuator"
	"github.com/adrianamaiaterosendo/projeto-dio-planejamento-financeiro.git/handler/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Init é onde eu tenho as chamadas http
func Init() {

	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transsaction/create", transaction.CreatedTransactions)
	http.HandleFunc("/health", actuator.Health)
	//cria metricas da aplicação onde podemos customizar
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
