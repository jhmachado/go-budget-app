package transaction

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/transactions", GetTransactions)
	http.HandleFunc("/transactions/create", CreateTransaction)

	http.ListenAndServe(":9000", nil)
}
