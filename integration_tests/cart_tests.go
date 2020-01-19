package integration_tests

import (
	"fmt"
	"net/http"

	"github.com/DATA-DOG/godog"
)

func (a *apiFeature) iAmUsingCartService() error {
	a.serverURI = "http://localhost:6379"
	return nil
}

func (a *apiFeature) iCallAddItemtoCartCart() error {
	endpoint := "/addcartitem/1"
	var response interface{}
	rc, err := SendHttp(http.MethodGet, endpoint, a.serverURI, nil, &response, getServiceHeaders())
	a.err = err.Error()
	a.respCode = rc
	return nil
}

func (a *apiFeature) iShouldReceiveSucessResponse() error {
	if a.respCode != 204 {
		return fmt.Errorf("Invalid Response, Items Not Added to Cart")
	}
	return nil
}

func (a *apiFeature) iCallGetCartItems() error {
	endpoint := "/getcartitems"
	rc, err := SendHttp(http.MethodGet, endpoint, a.serverURI, nil, &a.cartResponse, getServiceHeaders())
	a.err = err.Error()
	a.respCode = rc
	return nil
}

func (a *apiFeature) iShouldReceiveAddedItem() error {
	if a.cartResponse[0] != 1 {
		return fmt.Errorf("Added Item Cannot be found in the Cart")
	}
	return nil
}

func (a *apiFeature) iCallClearCart() error {
	endpoint := "/clearcart"
	var response interface{}
	rc, err := SendHttp(http.MethodDelete, endpoint, a.serverURI, nil, &response, getServiceHeaders())
	a.err = err.Error()
	a.respCode = rc
	return nil
}

func (a *apiFeature) iShouldReceiveEmptyCart() error {
	if len(a.cartResponse) != 0 {
		return fmt.Errorf("Cart Is Not Empty")
	}
	return nil
}

func brandFeatureContext(s *godog.Suite, a *apiFeature) {
	s.Step(`^I am using Cart Service$`, a.iAmUsingCartService)
	s.Step(`^I call add Cart Item$`, a.iCallAddItemtoCartCart)
	s.Step(`^I should receive response of success$`, a.iShouldReceiveSucessResponse)
	s.Step(`^I call get CartItems$`, a.iCallGetCartItems)
	s.Step(`^I should received the added Item from cart$`, a.iShouldReceiveAddedItem)
	s.Step(`^I call Clear Cart$`, a.iCallClearCart)
	s.Step(`^I should received empty cart$`, a.iShouldReceiveEmptyCart)

	
}
