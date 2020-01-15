package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/womblebob/uuid"
	"google.golang.org/grpc"

	item_service_pb "github.com/dilrandi/golang-practical-demo-shopping-cart/protos/itempb"
)

func main() {
	log.Infoln("Starting the Cart Service")
	defer log.Warningln("Exiting Cart service")
	cs := newCartService()
	cs.httpRouting()
}

func (cs *cartService) httpRouting() {
	log.Infoln("Starting the HTTP serving for Cart Servicing")

	router := httprouter.New()
	router.DELETE("/clearcart", httpErrorHandler(cs.clearCart))
	router.POST("/addcartitem/:itemid", httpErrorHandler(cs.AddItem))
	router.GET("/getcartitems", httpErrorHandler(cs.GetCartItems))

	log.Fatal(http.ListenAndServe(":8090", router))
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

//clearCart Clear Cart Items
func (cs *cartService) clearCart(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) (int, error) {
	log.Infof("Invoke Clear Cart.")
	defer log.Info("Exiting Clear Cart.")

	flush := cs.client.FlushAll()

	_, err := flush.Result()

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error Flushing Redis Data. error : %v", err)
	}

	return http.StatusOK, nil
}

//AddItem Add Item to the Cart
func (cs *cartService) AddItem(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) (int, error) {
	log.Infof("Invoke Add Item to Cart.")
	defer log.Info("Exiting Add Item to Cart.")

	stringid := parm.ByName("itemid")
	id, err := strconv.Atoi(stringid)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error Converting Id into Integer:%v", err)
	}

	exists, err := cs.validateItemGRPC(id)

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error validating Item GPRC:%v", err)
	}

	if !exists {
		return http.StatusInternalServerError, fmt.Errorf("Item with a item Id %s does not esists in Item service", stringid)
	}

	cart, err := cs.GetExistingCart()
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error when retriving cart information: %v", stringid)
	}
	if cart == "" {
		cs.client.Set(string(uuid.NewRandom()), stringid, 0)
	} else {
		cs.client.Append(cart, ","+stringid)
	}

	return cs.GetCartItems(rw, r, parm)

}

// GetCartItems Get Cart Items
func (cs *cartService) GetCartItems(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) (int, error) {
	log.Infof("Invoke Get cart Items.")
	defer log.Info("Exiting Get cart Items.")

	cart, err := cs.GetExistingCart()

	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error when retriving cart information : %v", err)
	}

	cartItems, err := cs.client.Get(cart).Result()
	log.Infoln(err)

	if err != nil {
		log.Errorln("No Items Found for the Cart.")
	}

	items := strings.Split(cartItems, ",")
	log.Infoln(items)
	json.NewEncoder(rw).Encode(items)

	return http.StatusOK, nil
}

//GetExistingCart get Existing Cart
func (cs *cartService) GetExistingCart() (string, error) {
	val, err := cs.client.Keys("*").Result()
	if err != nil {
		return "", err
	}

	if len(val) == 0 {
		return "", nil
	}

	return val[0], nil
}

// initRedisConnection Initializing a Redis Connection
func newStorage() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("Redis_Endpoint"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	s := client.Ping()
	r, err := s.Result()

	if err != nil {
		return nil, fmt.Errorf("Error while pinging the endpoint redis client, error : %v", err)

	}

	log.Info("Initializing a Redis Connection successful, with result : %s", r)
	return client, nil
}

func (cs *cartService) validateItemGRPC(itemID int) (bool, error) {
	con, err := grpc.Dial(fmt.Sprintf("%s:50051", cs.itemGRPCEndpoint), grpc.WithInsecure())

	if err != nil {
		return false, fmt.Errorf("Unable to connect to item service GRPC endpoint, error : %v", err)
	}

	defer con.Close()
	c := item_service_pb.NewItemServiceClient(con)
	req := &item_service_pb.IsItemExistsRequest{
		ItemID: int32(itemID),
	}
	res, err := c.IsItemExists(context.Background(), req)

	if err != nil {
		return false, fmt.Errorf("Unable to retrive the item from item service grpc endpoint, error: %v", err)
	}

	return res.Exists, nil
}

type cartService struct {
	client           *redis.Client
	itemGRPCEndpoint string
}

func newCartService() *cartService {
	client, err := newStorage()

	if err != nil {
		log.Fatalf("Unable to create redis client : %v", err)
	}

	itemGRPCEndpoint := os.Getenv("ITEM_GRPC_EP")

	return &cartService{
		client:           client,
		itemGRPCEndpoint: itemGRPCEndpoint,
	}

}
