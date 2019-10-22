package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Item struct {
	ID             int     `json:"id"`
	ItemCode       string  `json:"item_code"`
	Description    string  `json:"description"`
	UnitPrice      float32 `json:"unitPrice`
	PackingDetails PackingDetail
}

type PackingDetail struct {
	Width  float32 `json:"width`
	Height int     `json:"height`
	Weight int     `json:"weight`
}

var items []Item
var packingDetails []PackingDetail

func init() {
	packingDetails = []PackingDetail{
		PackingDetail{
			20.5,
			10,
			5,
		},
		PackingDetail{
			30.5,
			40,
			6,
		},
		PackingDetail{
			30.5,
			70,
			9,
		},
		PackingDetail{
			50.5,
			30,
			8,
		},
	}

	items = []Item{
		Item{
			1,
			"Itm001",
			"Tofee",
			400,
			packingDetails[0],
		},
		Item{
			2,
			"Itm002",
			"Choclate",
			200,
			packingDetails[1],
		},
		Item{
			3,
			"Itm003",
			"Buscuit",
			300,
			packingDetails[2],
		},
		Item{
			4,
			"Itm004",
			"Banana",
			400,
			packingDetails[3],
		},
	}

}

func main() {
	log.Println("Starting Item Service")
	http.HandleFunc("/getitem", GetItems)
	http.HandleFunc("/getitem/", GetItemById)
	http.ListenAndServe(":8080", nil)
	log.Println("Item Service Served")
}

//Get Items endpoint
func GetItems(rw http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(rw, r.URL.Path, http.StatusNotFound)
		return
	}
	log.Println("Get All Item Method Invoked")
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
	log.Println("Get Item By Id  Method Invoked")
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
