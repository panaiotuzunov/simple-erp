package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"simple-erp/internal/database"
	"strconv"

	"github.com/shopspring/decimal"
)

var grainTypes = map[string]struct{}{
	"wheat":     {},
	"barley":    {},
	"corn":      {},
	"sunflower": {},
	"canola":    {},
}

func (cfg *apiConfig) HandlerCreateEntranceReceipt(writer http.ResponseWriter, req *http.Request) {
	var requestData struct {
		TruckReg   string          `json:"truck_reg"`
		TrailerReg string          `json:"trailer_reg"`
		Gross      decimal.Decimal `json:"gross"`
		Tare       decimal.Decimal `json:"tare"`
		GrainType  string          `json:"grain_type"`
	}
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
	receipt, err := cfg.db.CreateEntranceReceipt(req.Context(), database.CreateEntranceReceiptParams{
		TruckReg:   requestData.TruckReg,
		TrailerReg: requestData.TrailerReg,
		Gross:      requestData.Gross,
		Tare:       requestData.Tare,
		GrainType:  requestData.GrainType,
	})
	if err != nil {
		log.Printf("Error creating entrance receipt: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error creating entrance receipt")
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
		writeErrorResponse(writer, http.StatusBadRequest, "error converting id string to int")
		return
	}
	receipt, err := cfg.db.GetEntranceReceiptByID(req.Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			writeErrorResponse(writer, http.StatusNotFound, "Not found")
			return
		}
		writeErrorResponse(writer, http.StatusInternalServerError, "DB error")
		return
	}
	writeJSONResponse(writer, http.StatusOK, receipt)
}
