package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type StockItem struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Warehouse   string `json:"warehouse"`
}

var stockData = map[string]StockItem{
	"PROD-001": {ProductID: "PROD-001", ProductName: "Widget A", Quantity: 150, Warehouse: "WH-EAST"},
	"PROD-002": {ProductID: "PROD-002", ProductName: "Widget B", Quantity: 0, Warehouse: "WH-WEST"},
	"PROD-003": {ProductID: "PROD-003", ProductName: "Gadget C", Quantity: 42, Warehouse: "WH-EAST"},
}

func handleGetStock(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("product_id")
	if productID == "" {
		http.Error(w, `{"error":"product_id required"}`, http.StatusBadRequest)
		return
	}
	stock, ok := stockData[productID]
	if !ok {
		http.Error(w, `{"error":"product not found"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok","service":"inventory-service"}`))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	http.HandleFunc("/stock", handleGetStock)
	http.HandleFunc("/health", handleHealth)
	log.Printf("inventory-service listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
