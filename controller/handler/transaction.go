package handler

import (
	"encoding/json"
	"github.com/leopedroso45/UpnidChallenge/domain"
	"io/ioutil"
	"net/http"
)

/*GetTransactions is the "/v1.0/transactions" handler that returns a json response if no error occurs*/
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	var transactionsJson []domain.Transaction
	transactionsBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error trying to read the request body, check if everything is correct.", 400)
	}
	err2 := json.Unmarshal(transactionsBytes, &transactionsJson)
	if err2 != nil {
		http.Error(w, "Error trying to read the request body, check if everything is correct.\n"+err2.Error(), 400)
	}
	resultSlice := domain.CheckFraud(transactionsJson)
	resultJson, err3 := json.Marshal(resultSlice)
	if err3 != nil {
		http.Error(w, err3.Error(), 500)
	}

	if err2 == nil && err3 == nil {
		_, err = w.Write(resultJson)
	}
}
