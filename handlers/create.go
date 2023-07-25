package handlers

import (
	"challenge/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CurrencyConverter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	conversion, err := models.ExchangeHandlerService(params)
	if err != nil {
		http.Error(w, "Failed to process conversion request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conversion); err != nil {
		http.Error(w, "Failed to encode conversion response", http.StatusInternalServerError)
	}
}
