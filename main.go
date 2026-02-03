package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"simple-erp/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Opening DB connection failed")
	}
	cfg := apiConfig{
		db: database.New(db),
	}
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	// home page
	mux.Handle("/", http.FileServer(http.Dir("./html")))
	// etrance-receipts endpoint
	mux.HandleFunc("POST /api/entrance-receipts/", cfg.HandlerCreateEntranceReceipt)
	mux.HandleFunc("GET /api/entrance-receipts/", cfg.HandlerGetAllEntranceReceipts)
	mux.HandleFunc("GET /api/entrance-receipts/{receiptID}", cfg.HandlerGetEntranceReceiptByID)
	// exit-receipts endpoint
	mux.HandleFunc("POST /api/exit-receipts/", cfg.HandlerCreateExitReceipt)
	mux.HandleFunc("GET /api/exit-receipts/", cfg.HandlerGetAllExitReceipts)
	mux.HandleFunc("GET /api/exit-receipts/{receiptID}", cfg.HandlerGetExitReceiptByID)
	// purchases endpoint
	mux.HandleFunc("POST /api/purchases/", cfg.HandlerCreatePurchase)
	mux.HandleFunc("GET /api/purchases/", cfg.HandlerGetAllPurchases)
	// sales endpoint
	mux.HandleFunc("POST /api/sales/", cfg.HandlerCreateSale)
	mux.HandleFunc("GET /api/sales/", cfg.HandlerGetAllSales)
	// transports endpoint
	mux.HandleFunc("POST /api/transports/", cfg.HandlerCreateTransport)
	mux.HandleFunc("GET /api/transports/", cfg.HandlerGetAllTransports)
	// companies endpoint
	mux.HandleFunc("POST /api/companies/", cfg.HandlerCreateCompany)
	mux.HandleFunc("GET /api/companies/", cfg.HandlerGetAllCompanies)
	// reports endpoint
	mux.HandleFunc("GET /api/reports/movements/", cfg.HandlerReportMovements)
	mux.HandleFunc("GET /api/reports/inventory/", cfg.HandlerReportInventory)

	log.Print("Server is running")
	server.ListenAndServe()
}
