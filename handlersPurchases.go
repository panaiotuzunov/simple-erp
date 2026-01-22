package main

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-erp/internal/database"
)

func (cfg *apiConfig) HandlerCreatePurchase(writer http.ResponseWriter, req *http.Request) {
	var requestData database.CreatePurchaseParams
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&requestData); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		writeErrorResponse(writer, http.StatusBadRequest, "Error decoding JSON")
		return
	}
	purchase, err := cfg.db.CreatePurchase(req.Context(), requestData)
	if err != nil {
		log.Printf("Error creating purchase: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error creating purchase")
		return
	}
	writeJSONResponse(writer, http.StatusCreated, purchase)
}
