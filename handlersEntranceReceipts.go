package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simple-erp/internal/database"
	"strconv"
)

func (cfg *apiConfig) HandlerCreateEntranceReceipt(writer http.ResponseWriter, req *http.Request) {
	var requestData database.CreateEntranceReceiptParams
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
	if _, err := cfg.db.GetPurchaseById(req.Context(), requestData.PurchaseID); err != nil {
		if err == sql.ErrNoRows {
			writeErrorResponse(writer, http.StatusNotFound, "Purchase not found")
			return
		}
		log.Printf("Error retrieving purchase: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error retrieving purchase")
		return
	}
	receipt, err := cfg.db.CreateEntranceReceipt(req.Context(), requestData)
	if err != nil {
		log.Printf("Error creating receipt: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error creating receipt")
		return
	}
	writeJSONResponse(writer, http.StatusCreated, receipt)
}

func (cfg *apiConfig) HandlerGetAllEntranceReceipts(writer http.ResponseWriter, req *http.Request) {
	receipts, err := cfg.db.GetAllEntranceReceipts(req.Context())
	if err != nil {
		log.Printf("DB error - %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "DB error")
		return
	}
	if len(receipts) == 0 {
		writeJSONResponse(writer, http.StatusOK, []database.EntranceReceipt{})
		return
	}
	writeJSONResponse(writer, http.StatusOK, receipts)
}

func (cfg *apiConfig) HandlerGetEntranceReceiptByID(writer http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("receiptID"))
	if err != nil {
		log.Printf("Error converting id str to int: %v", err)
		writeErrorResponse(writer, http.StatusBadRequest, "error converting id string to int")
		return
	}
	receipt, err := cfg.db.GetEntranceReceiptByID(req.Context(), int32(id))
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
