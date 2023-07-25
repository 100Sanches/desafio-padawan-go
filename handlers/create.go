package handlers

import (
	"challenge/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CurrencyConverter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	conversion, err := models.ExchangeHandlerService(params)
	if err != nil {
		log.Printf("Failed to process conversion request: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(conversion); err != nil {
		http.Error(w, "Failed to encode conversion response", http.StatusInternalServerError)
	}
}
