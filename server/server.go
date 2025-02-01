package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Struct for JSON responses
type Account struct {
	Owner   string  `json:"owner"`
	Balance float32 `json:"balance"`
	Id      string  `json:"id"`
}

var accounts = map[string]*Account{
	"1": {Owner: "Alice", Balance: 100.0, Id: "1"},
	"2": {Owner: "Bob", Balance: 50.0, Id: "2"},
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the URL path
	id := strings.TrimPrefix(r.URL.Path, "/balance/")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	account, exists := accounts[id]
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id     string  `json:"id"`
		Amount float32 `json:"amount"`
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "Deposit must be positive", http.StatusBadRequest)
		return
	}

	account, exists := accounts[req.Id]
	if !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	account.Balance += req.Amount

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func RunServer() {
	// By adding the GET here, the server will fail any request to this handler of any other method
	// The trailing slash is necessary because here we are using the id as a path parameter
	http.HandleFunc("GET /balance/", balanceHandler)
	http.HandleFunc("/deposit", depositHandler)
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}
