package integration_tests

import ("github.com/DATA-DOG/godog"
"fmt"
"net/http"
)

//PackingDetail ..
type PackingDetailResponse struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

type ItemResponse struct {
	ID             int           `json:"id"`
	ItemCode       string        `json:"item_code"`
	Description    string        `json:"description"`
	UnitPrice      float32       `json:"unit_price"`
	PackingDetails PackingDetailResponse `json:"packing_details"`
}

type cartFeature struct {
	serverURI string
	
	respCode int
	err string

	itemResponse ItemResponse 
	cartResponse []string
}

func (c *cartFeature) iAmUsingCartService() error {
	c.serverURI = "http://cart_service:8090"
	return nil
}

func (c *cartFeature) iCallAddCartItem() error {
	endpoint := "/addcartitem/1"
	var response interface{}
	rc, err := SendHttp(http.MethodPost, endpoint, c.serverURI, nil, &response, getServiceHeaders())
	if err!= nil{
		return fmt.Errorf("Item Not Added to Cart :%v",err)
	}
	c.respCode = rc
	return nil
}

func (c *cartFeature) iCallGetCartItems() error {
	endpoint := "/getcartitems"
	rc, err := SendHttp(http.MethodGet, endpoint, c.serverURI, nil, &c.cartResponse, getServiceHeaders())
	if err!= nil{
		return fmt.Errorf("Item Not Added to Cart :%v",err)
	}
	
	c.respCode = rc
	return nil
}

func (c *cartFeature) iShouldReceiveResponseOfSuccess() error {
	if c.respCode != 200 {
		return fmt.Errorf("Invalid Response, Items Not Added to Cart:%d",c.respCode)
	}
	return nil
}

func (c *cartFeature) iShouldReceivedTheAddedItemFromCart() error {
	
	if c.cartResponse[0] != "1" {
		return fmt.Errorf("Added Item Cannot be found in the Cart")
	}
	return nil
}

func (c *cartFeature) iShouldReceiveEmptyCart() error {
	if c.cartResponse[0] == "1" {
		return fmt.Errorf("Cart Is Not Empty")
	}
	return nil
}

func (c *cartFeature) iCallClearCart() error {
	endpoint := "/clearcart"
	var response interface{}
	rc, err := SendHttp(http.MethodDelete, endpoint, c.serverURI, nil, &response, getServiceHeaders())
	if err !=nil{
		return fmt.Errorf("Method Delete Error :%v",err)
	}
	c.respCode = rc
	return nil
}

func CartFeatureContext(s *godog.Suite) {
	cs := new(cartFeature)
	s.Step(`^I am using Cart Service$`, cs.iAmUsingCartService)
	s.Step(`^I call add Cart Item$`, cs.iCallAddCartItem)
	s.Step(`^I call Clear Cart$`, cs.iCallClearCart)
	s.Step(`^I call get CartItems$`, cs.iCallGetCartItems)
	s.Step(`^I should receive response of success$`, cs.iShouldReceiveResponseOfSuccess)
	s.Step(`^I should received the added Item from cart$`, cs.iShouldReceivedTheAddedItemFromCart)
	s.Step(`^I should receive empty cart$`, cs.iShouldReceiveEmptyCart)

}
