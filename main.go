package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	inventory "inventory/Inventory"
	"log"
	"net/http"
)

var backpack = inventory.Inventory{}

type SearchRequest struct {
	Query string `json:"query"`
}

type SearchResponse struct {
	Results []string `json:"results"`
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/backpack", handleBackpack)
	http.HandleFunc("/search", handleSearch)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	slot1 := inventory.InventorySlot{
		ItemName:   "coins",
		StackLimit: 100,
		StackSize:  50,
		ImagePath:  "goldCoin.png",
		ItemWeight: 0.1,
		Category:   "money",
	}

	backpack = inventory.Inventory{
		WeightLimit: 100,
		Name:        "Backpack",
		ItemArray:   []inventory.InventorySlot{},
	}
	backpack.AddItem(slot1)
	backpack.AddItem(inventory.InventorySlot{})
	fmt.Println(backpack)

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleBackpack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(backpack.ItemArray)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Println(req.Query)
	json.NewEncoder(w).Encode(map[string]string{"message": "Search functionality not implemented yet"})
}
