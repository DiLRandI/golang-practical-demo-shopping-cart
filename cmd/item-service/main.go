package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Item struct {
	ID          int     `json:"id"`
	ItemCode    string  `json:"item_code"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unitPrice`
	Width       float32 `json:"width`
	Height      int     `json:"height`
	Weight      int     `json:"weight`
}

var items []Item

func init() {
	items = []Item{
		Item{
			1,
			"Itm001",
			"Tofee",
			100,
			20.5,
			10,
			5,
		},
		Item{
			2,
			"Itm002",
			"Choclate",
			200,
			30.5,
			40,
			6,
		},
		Item{
			3,
			"Itm003",
			"Buscuit",
			300,
			30.5,
			40,
			7,
		},
		Item{
			4,
			"Itm004",
			"Banana",
			400,
			30.5,
			40,
			7,
		},
	}
}

func main() {
	fmt.Println("Hello from item service.")
	http.HandleFunc("/getitem", GetItems)
	http.HandleFunc("/getitem/", GetItemById)
	http.ListenAndServe(":8080", nil)

}

//Get Items endpoint
func GetItems(rw http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(rw, r.URL.Path, http.StatusNotFound)
		return
	}

	json.NewEncoder(rw).Encode(items)
}

//Get Item By Id endpoint
func GetItemById(rw http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(rw, "Invalid Method", http.StatusNotFound)
		return
	}

	stringId := strings.Replace(r.URL.Path, "/getitem/", "", 1)

	id, err := strconv.Atoi(stringId)

	if err != nil {
		http.Error(rw, "Invalid Type of ItemId", http.StatusBadRequest)
		return
	}

	index := indexOf(id)

	if index == -1 {
		http.Error(rw, "Item Not Found", http.StatusNotFound)
		return
	}
	json.NewEncoder(rw).Encode(items[index])

}

func indexOf(element int) int {
	for k, v := range items {
		if element == v.ID {
			return k
		}
	}
	return -1 //not found.
}
