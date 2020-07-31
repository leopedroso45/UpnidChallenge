package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/leopedroso45/UpnidChallenge/controller/handler"
	"log"
	"net/http"
)

func main() {

	log.Println("Server running...")

	router := mux.NewRouter()
	router.HandleFunc("/v1.0/transactions", handler.GetTransactions).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(router)))

}
