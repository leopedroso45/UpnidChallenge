package handler

import (
	"encoding/json"
	"github.com/leopedroso45/UpnidChallenge/controller"
	"github.com/leopedroso45/UpnidChallenge/domain"
	"io/ioutil"
	"net/http"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	var transactionsJson []domain.Transaction
	transactionsBytes, err := ioutil.ReadAll(r.Body)
	controller.CheckError(err)
	err = json.Unmarshal(transactionsBytes, &transactionsJson)
	controller.CheckError(err)
	resultSlice := domain.CheckFraud(transactionsJson)
	resultJson, err := json.Marshal(resultSlice)
	controller.CheckError(err)

	_, err = w.Write(resultJson)
}
