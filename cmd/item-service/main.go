package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
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
	router := httprouter.New()
	router.GET("/getitem", GetItems)
	router.GET("/getitem/:ID", GetItemByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}

//GetItems endpoint
func GetItems(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Info("GetAllItem Method Invoked")
	defer log.Info("GetAllItem  Exiting")

	json.NewEncoder(rw).Encode(items)
}

//GetItemByID Id endpoint
func GetItemByID(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) {
	log.Info("GetItemByID  Method Invoked")
	defer log.Info("GetItemByID Exiting")

	stringID := parm.ByName("ID")
	id, err := strconv.Atoi(stringID)

	if err != nil {
		log.Errorf("error while reading ID param, %v", err)
		http.Error(rw, "Invalid Type of ItemId", http.StatusBadRequest)
		return
	}

	index := indexOf(id)

	if index == -1 {
		log.Warningf("No data found to given ID : %v", id)
		http.Error(rw, "", http.StatusNotFound)
		return
	}

	log.Printf("Response value, %+v", items[index])
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
