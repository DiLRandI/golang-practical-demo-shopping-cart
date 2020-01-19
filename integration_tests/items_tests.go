package integration_tests

import (
	"fmt"
	"net/http"

	"github.com/DATA-DOG/godog"
)

func (a *apiFeature) iAmUsingItemService() error {
	a.serverURI = "http://localhost:8080"
	return nil
}

func (a *apiFeature) iCallGetEndpointInItems() error {
	endpoint := "/getitem"
	var response interface{}
	rc, err := SendHttp(http.MethodGet, endpoint, a.serverURI, nil, &response, getServiceHeaders())
	a.err = err.Error()
	a.respCode = rc
	return nil
}

func (a *apiFeature) iShouldReceiveItemList() error {
	if a.respCode != 204 {
		return fmt.Errorf("Invalid Response, Items Not Received.")
	}
	return nil
}

func (a *apiFeature) iCallGetItemEndpoint() error {
	endpoint := "/getitem/1"
	
	rc, err := SendHttp(http.MethodPost, endpoint, a.cartServerUri, nil, &a.itemResponse, getServiceHeaders())
	a.err = err.Error()
	a.respCode = rc
	return nil
}

func (a *apiFeature) iShouldReceiveItem() error {
	if a.itemResponse.ID != 1 {
		return fmt.Errorf("Invalid Item Id")
	}
	return nil
}

func itemFeatureContext(s *godog.Suite, a *apiFeature) {
	s.Step(`^I am using Item Service$`, a.iAmUsingItemService)
	s.Step(`^I call Get endpoint in Items$`, a.iCallGetEndpointInItems)
	s.Step(`^I should receive Items$`, a.iShouldReceiveItemList)
	s.Step(`^I call Get endpoint in Item$`, a.iCallGetItemEndpoint)
	s.Step(`^I should receive an Item$`, a.iShouldReceiveItem)
	
}
