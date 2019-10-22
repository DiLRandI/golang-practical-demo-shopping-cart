package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//Item ...
type Item struct {
	ID             int           `json:"id"`
	ItemCode       string        `json:"item_code"`
	Description    string        `json:"description"`
	UnitPrice      float32       `json:"unit_price"`
	PackingDetails PackingDetail `json:"packing_details"`
}

//PackingDetail ..
type PackingDetail struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

var items []Item
var packingDetails []PackingDetail

func init() {
	packingDetails = []PackingDetail{
		PackingDetail{
			20,
			10,
			5,
		},
		PackingDetail{
			30,
			40,
			6,
		},
		PackingDetail{
			30,
			70,
			9,
		},
		PackingDetail{
			50,
			30,
			8,
		},
	}

	items = []Item{
		Item{
			1,
			"Itm001",
			"Toffee",
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
			"Biscuit",
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
	// log.Println("Starting Item Service")
	// http.HandleFunc("/getitem", GetItems)
	// http.HandleFunc("/getitem/", GetItemByID)
	// http.ListenAndServe(":8080", nil)
	// log.Println("Item Service Served")

	router := httprouter.New()
	router.GET("/getitem", GetItems)
	router.GET("/getitem/:ID", GetItemByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//GetItems endpoint
func GetItems(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("GetAllItem Method Invoked")
	defer log.Println("GetAllItem  Exiting")

	json.NewEncoder(rw).Encode(items)
}

//GetItemByID Id endpoint
func GetItemByID(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) {
	log.Println("GetItemByID  Method Invoked")
	defer log.Println("GetItemByID Exiting")

	stringID := parm.ByName("ID")

	id, err := strconv.Atoi(stringID)

	if err != nil {
		http.Error(rw, "Invalid Type of ItemId", http.StatusBadRequest)
		return
	}

	index := indexOf(id)

	if index == -1 {
		http.Error(rw, "", http.StatusNotFound)
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
