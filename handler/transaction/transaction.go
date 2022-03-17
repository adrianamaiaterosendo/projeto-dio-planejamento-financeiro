package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adrianamaiaterosendo/projeto-dio-planejamento-financeiro.git/model/transaction"
	"github.com/adrianamaiaterosendo/projeto-dio-planejamento-financeiro.git/utils"
)

//GetTransactions - Traz as informações das transações
func GetTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	dateCreate, err := utils.StringToTime("2006-01-02T15:04:05")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.00,
			Type:      0,
			CreatedAt: dateCreate,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)

}

//CreatedTransactions - Cadastra uma transação
func CreatedTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)

}
