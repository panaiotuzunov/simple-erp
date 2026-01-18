package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, text string) {
	writeJSONResponse(w, statusCode, errorResponse{Error: text})
}
