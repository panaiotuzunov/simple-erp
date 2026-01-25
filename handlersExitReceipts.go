package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simple-erp/internal/database"
	"strconv"
)

func (cfg *apiConfig) HandlerCreateExitReceipt(writer http.ResponseWriter, req *http.Request) {
	var requestData database.CreateExitReceiptParams
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&requestData); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		writeErrorResponse(writer, http.StatusBadRequest, "Error decoding JSON")
		return
	}
	if _, exist := grainTypes[requestData.GrainType]; !exist {
		writeErrorResponse(writer, http.StatusBadRequest, "Grain type not found")
		return
	}
	if _, err := cfg.db.GetSaleById(req.Context(), requestData.SaleID); err != nil {
		if err == sql.ErrNoRows {
			writeErrorResponse(writer, http.StatusNotFound, "Sale not found")
			return
		}
		log.Printf("Error retrieving sale: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error retrieving sale")
		return
	}
	receipt, err := cfg.db.CreateExitReceipt(req.Context(), requestData)
	if err != nil {
		log.Printf("Error creating receipt: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error creating receipt")
		return
	}
	writeJSONResponse(writer, http.StatusCreated, receipt)
}

func (cfg *apiConfig) HandlerGetAllExitReceipts(writer http.ResponseWriter, req *http.Request) {
	receipts, err := cfg.db.GetAllExitReceipts(req.Context())
	if err != nil {
		log.Printf("DB error - %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "DB error")
		return
	}
	if len(receipts) == 0 {
		writeJSONResponse(writer, http.StatusOK, []database.ExitReceipt{})
		return
	}
	writeJSONResponse(writer, http.StatusOK, receipts)
}

func (cfg *apiConfig) HandlerGetExitReceiptByID(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("receiptID"))
	if err != nil {
		log.Printf("Error converting id str to int: %v", err)
		writeErrorResponse(writer, http.StatusBadRequest, "error converting id string to int")
		return
	}
	receipt, err := cfg.db.GetExitReceiptByID(req.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			writeErrorResponse(writer, http.StatusNotFound, "Not found")
			return
		}
		log.Printf("Error fetching receipt by id: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "DB error")
		return
	}
	writeJSONResponse(writer, http.StatusOK, receipt)
}
