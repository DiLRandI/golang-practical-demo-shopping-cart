package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	router.GET("/getitem", httpErrorHandler(GetItems))
	router.GET("/getitem/:ID", httpErrorHandler(GetItemByID))

	log.Fatal(http.ListenAndServe(":8080", router))
}

// httpErrorHandler is a wrapper for http handler.
// this will handle errors return from the handlers.
func httpErrorHandler(hf func(http.ResponseWriter, *http.Request, httprouter.Params) (int, error)) func(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) {
	return func(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) {
		if s, err := hf(rw, r, parm); err != nil {
			log.Errorf("Error return from the handler , error : %v ", err)
			http.Error(rw, err.Error(), s)
		}
	}
}

//GetItems endpoint
func GetItems(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) (int, error) {
	log.Info("GetAllItem Method Invoked")
	defer log.Info("GetAllItem  Exiting")

	json.NewEncoder(rw).Encode(items)
	return http.StatusOK, nil
}

//GetItemByID Id endpoint
func GetItemByID(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) (int, error) {
	log.Info("GetItemByID  Method Invoked")
	defer log.Info("GetItemByID Exiting")

	stringID := parm.ByName("ID")
	id, err := strconv.Atoi(stringID)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("error while reading ID param, %v", err)
	}

	index := indexOf(id)

	if index == -1 {
		return http.StatusInternalServerError, fmt.Errorf("No data found to given ID : %v", id)
	}

	log.Printf("Response value, %+v", items[index])
	json.NewEncoder(rw).Encode(items[index])
	return http.StatusOK, nil
}

func indexOf(element int) int {
	for k, v := range items {
		if element == v.ID {
			return k
		}
	}
	return -1 //not found.
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
