package main

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-erp/internal/database"

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
		log.Printf("Error decoding JSON: %s", err)
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
		log.Printf("Error creating entrance receipt: %s", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error creating entrance receipt")
		return
	}
	writeJSONResponse(writer, http.StatusCreated, receipt)
}
