package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/tickets/{username}", GetTicketsByUsernameHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/tickets/{uid}", CancelTicketByUidHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/tickets", BuyTicketHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/manage/health", GetHealth).Methods("GET")

	return router
}
