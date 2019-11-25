package main

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"google.golang.org/grpc"

	item_service_pb "github.com/dilrandi/golang-practical-demo-shopping-cart/protos/itempb"
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

func main() {
	log.Infoln("Starting the Item Service")
	defer log.Warningln("Exiting item service")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go httpRouting()
	go grpcRouting()

	<-sig
}

func httpRouting() {
	log.Infoln("Starting the HTTP serving")

	router := httprouter.New()
	router.GET("/getitem", GetItems)
	router.GET("/getitem/:ID", GetItemByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func grpcRouting() {
	log.Infoln("Starting the GRPC serving")

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Unable to create listener, error : %v", err)
	}

	s := grpc.NewServer()

	item_service_pb.RegisterItemServiceServer(s, &itemGrpc{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Unable to start GRPC server, error : %v", err)
	}
}

type itemGrpc struct {
}

func (*itemGrpc) IsItemExists(ctx context.Context, req *item_service_pb.IsItemExistsRequest) (*item_service_pb.IsItemExistsResponse, error) {
	log.Infof("Invoke GRPC endpoint : IsItemExists with ID : %v", req.ItemID)
	defer log.Info("Exiting IsItemExists GRPC handler.")
	index := indexOf(int(req.ItemID))
	res := new(item_service_pb.IsItemExistsResponse)

	if index != -1 {
		log.Infof("Item with ID: %v, found", req.ItemID)
		res.Exists = true
		return res, nil
	}

	log.Warningf("Item with ID: %v, not exists", req.ItemID)
	res.Exists = false
	return res, nil
}
