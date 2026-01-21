package main

import (
	"log"
	"net/http"
	"simple-erp/internal/database"
)

func (cfg *apiConfig) HandlerReportMovements(writer http.ResponseWriter, req *http.Request) {
	receipts, err := cfg.db.GetAllReceiptsUnion(req.Context())
	if err != nil {
		log.Printf("DB error - %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "DB error")
		return
	}
	if len(receipts) == 0 {
		writeJSONResponse(writer, http.StatusOK, []database.GetAllReceiptsUnionRow{})
		return
	}
	writeJSONResponse(writer, http.StatusOK, receipts)
}
