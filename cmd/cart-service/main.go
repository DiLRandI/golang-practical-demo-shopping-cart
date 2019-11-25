package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/womblebob/uuid"
)

var client *redis.Client
var a uuid.UUID

func main() {
	log.Infoln("Starting the Cart Service")
	defer log.Warningln("Exiting Cart service")

	httpRouting()
}

func httpRouting() {
	log.Infoln("Starting the HTTP serving for Cart Servicing")

	router := httprouter.New()
	router.DELETE("/Clear", ClearCart)
	router.POST("/AddItem/:ItemId", AddItem)
	router.GET("/GetItems", GetCartItems)

	log.Fatal(http.ListenAndServe(":8090", router))
}

//ClearCart Clear Cart Items
func ClearCart(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) {
	InitRedisConnection()
	flush := client.FlushAll()
	if flush != nil {
		log.Infoln(flush)
	}
}

//AddItem Add Item to the Cart
func AddItem(rw http.ResponseWriter, r *http.Request, parm httprouter.Params) {
	InitRedisConnection()
	id := parm.ByName("ItemId")
	cart, err := GetExistingCart()
	if err != nil {
		log.Errorln(err)
		http.Error(rw, "Error when Checking for Existing Carts", http.StatusBadRequest)
	}
	if cart == "" {
		client.Set(string(uuid.NewRandom()), id, 0)
	} else {
		client.Append(cart, ","+id)
	}

}

// GetCartItems Get Cart Items
func GetCartItems(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	InitRedisConnection()
	cart, err := GetExistingCart()
	if err != nil {
		log.Errorln(err)
		http.Error(rw, "Error when Checking for Existing Carts", http.StatusBadRequest)
	}
	cartItems, err := client.Get(cart).Result()
	if err != nil {
		log.Errorln(err)
		http.Error(rw, "Error when retriving Items", http.StatusBadRequest)
	}
	items := strings.Split(cartItems, ",")
	log.Infoln(items)
	json.NewEncoder(rw).Encode(items)
}

//GetExistingCart get Existing Cart
func GetExistingCart() (string, error) {
	val, err := client.Keys("*").Result()
	if err != nil {
		log.Error(err)
		return "", err
	}

	if len(val) == 0 {
		return "", nil
	}

	return val[0], nil
}

//InitRedisConnection Initializing a Redis Connection
func InitRedisConnection() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Infoln("Initializing a Redis Connection.")
}
