package main

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-erp/internal/database"
)

func (cfg *apiConfig) HandlerCreateCompany(writer http.ResponseWriter, req *http.Request) {
	var requestData database.CreateCompanyParams
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&requestData); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		writeErrorResponse(writer, http.StatusBadRequest, "Error decoding JSON")
		return
	}
	company, err := cfg.db.CreateCompany(req.Context(), requestData)
	if err != nil {
		log.Printf("Error creating company: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error creating company")
		return
	}
	writeJSONResponse(writer, http.StatusCreated, company)
}

func (cfg *apiConfig) HandlerGetAllCompanies(writer http.ResponseWriter, req *http.Request) {
	companies, err := cfg.db.GetAllCompanies(req.Context())
	if err != nil {
		log.Printf("Error retrieving companies: %v", err)
		writeErrorResponse(writer, http.StatusInternalServerError, "Error retrieving companies")
		return
	}
	writeJSONResponse(writer, http.StatusOK, companies)
}