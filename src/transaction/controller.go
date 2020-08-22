package transaction

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/jhmachado/app/util"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var transactions = Transactions{
		Transaction{
			Title:     "Salário",
			Amount:    1220.0,
			Type: 	   0,
			CreatedAt: util.ConvertStringToTime("2020-08-22T13:43:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newTransaction = Transaction{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &newTransaction)
}
