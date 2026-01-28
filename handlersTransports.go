package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simple-erp/internal/database"
)

func (cfg *apiConfig) HandlerCreateTransport(writer http.ResponseWriter, req *http.Request) {
	var requestData database.CreateTransportParams
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&requestData); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		writeErrorResponse(writer, http.StatusBadRequest, "Error decoding JSON")
		return
	}
	purchase, err := cfg.db.GetPurchaseById(req.Context(), requestData.PurchaseID)
	if err != nil {
		if err == sql.ErrNoRows {
			writeErrorResponse(writer, http.StatusNotFound, "Purchase not found")
			return
		}
		log.Printf("Error retrieving purchase: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error retrieving purchase")
		return
	}
	sale, err := cfg.db.GetSaleById(req.Context(), requestData.SaleID)
	if err != nil {
		if err == sql.ErrNoRows {
			writeErrorResponse(writer, http.StatusNotFound, "Sale not found")
			return
		}
		log.Printf("Error retrieving sale: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error retrieving sale")
		return
	}
	if sale.GrainType != purchase.GrainType {
		writeErrorResponse(writer, http.StatusBadRequest, "grain_type mismatch. Purchase and sale grain_type must be the same")
		return
	}
	transport, err := cfg.db.CreateTransport(req.Context(), requestData)
	if err != nil {
		log.Printf("Error creating transport: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error creating transport")
		return
	}
	writeJSONResponse(writer, http.StatusOK, transport)
}
func (cfg *apiConfig) HandlerGetAllTransports(writer http.ResponseWriter, req *http.Request) {
	transports, err := cfg.db.GetAllTransports(req.Context())
	if err != nil {
		log.Printf("Error retrieving transports: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error retrieving transports")
		return
	}
	if len(transports) == 0 {
		writeJSONResponse(writer, http.StatusOK, []database.GetAllTransportsRow{})
		return
	}
	writeJSONResponse(writer, http.StatusOK, transports)
}
