package main

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-erp/internal/database"
)

func (cfg *apiConfig) HandlerCreateSale(writer http.ResponseWriter, req *http.Request) {
	var requestData database.CreateSaleParams
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&requestData); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		writeErrorResponse(writer, http.StatusBadRequest, "Error decoding JSON")
		return
	}
	sale, err := cfg.db.CreateSale(req.Context(), requestData)
	if err != nil {
		log.Printf("Error creating sale: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error creating sale")
		return
	}
	writeJSONResponse(writer, http.StatusCreated, sale)
}

func (cfg *apiConfig) HandlerGetAllSales(writer http.ResponseWriter, req *http.Request) {
	sales, err := cfg.db.GetAllSales(req.Context())
	if err != nil {
		log.Printf("Error retrieving sales: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error retrieving sales")
		return
	}
	if len(sales) == 0 {
		writeJSONResponse(writer, http.StatusOK, []database.Sale{})
		return
	}
	writeJSONResponse(writer, http.StatusOK, sales)
}
