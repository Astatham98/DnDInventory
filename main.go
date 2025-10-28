package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"html/template"
	inventory "inventory/Inventory"
	"log"
	"net/http"
	"os"
	"strings"
)

var backpack = inventory.Inventory{}
var items_csv = readCsvFile("Mythical_Items.csv")

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
	item_name := req.Query
	data_row := filterByName(item_name)
	if data_row == nil {
		fmt.Println("Could not find item in datarow")
		return
	}
	itemSlot := generateItemSlot(data_row)
	backpack.AddItem(itemSlot)
	json.NewEncoder(w).Encode("success")

	// fmt.Println(item_name)
	// json.NewEncoder(w).Encode(map[string]string{"message": "Search functionality not implemented yet"})
}

func readCsvFile(filepath string) [][]string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Could not open desired file", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("could not read records", err)
	}

	return records
}

func filterByName(item string) []string {
	for _, row := range items_csv {
		if strings.EqualFold(row[0], item) {
			return row
		}
	}
	return nil
}

func getItemStackSize(item string) int {
	switch item {
	case "ARMOR":
		return 2
	case "WEAPON":
		return 3
	case "WONDROUS_ITEMS":
		return 2
	default:
		return 0
	}
}
func generateItemSlot(row []string) inventory.InventorySlot {
	stackLimit := getItemStackSize(row[3])

	slot := inventory.InventorySlot{
		ItemName:   row[0],
		StackLimit: stackLimit,
		StackSize:  1,
		ImagePath:  row[1],
		ItemWeight: 1,
		Category:   row[3],
	}
	return slot
}
