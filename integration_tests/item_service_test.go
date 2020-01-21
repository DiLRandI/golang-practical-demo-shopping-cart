package integration_tests

import ("github.com/DATA-DOG/godog"
"fmt"
"net/http"
)

type itemFeature struct {
	serverURI string
	
	respCode int
	err string

	itemResponse ItemResponse 
	cartResponse []int
}

func (i *itemFeature) iAmUsingItemService() error {
	i.serverURI = "http://item_service:8080"
	return nil
}

func (i *itemFeature) iCallGetEndpointInItems() error {
	endpoint := "/getitem"
	var response interface{}
	rc, _ := SendHttp(http.MethodGet, endpoint, i.serverURI, nil, &response, getServiceHeaders())
	i.respCode = rc
	return nil
}

func (i *itemFeature) iShouldReceiveItems() error {
	if i.respCode != 200 {
		return fmt.Errorf("Invalid Response, Items Not received")
	}
	return nil
}

func (i *itemFeature) iCallGetEndpointInItem() error {
	endpoint := "/getitem/1"
	rc, _ := SendHttp(http.MethodGet, endpoint, i.serverURI, nil, &i.itemResponse, getServiceHeaders())
	i.respCode = rc
	return nil
}

func (i *itemFeature) iShouldReceiveAnItem() error {

	if i.itemResponse.ID != 1{
		return fmt.Errorf("Invalid Item Returned %v",i.itemResponse.ID)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	is := new(itemFeature)
	s.Step(`^I am using Item Service$`, is.iAmUsingItemService)
	s.Step(`^I call Get endpoint in Items$`, is.iCallGetEndpointInItems)
	s.Step(`^I should receive Items$`, is.iShouldReceiveItems)
	s.Step(`^I call Get endpoint in Item$`, is.iCallGetEndpointInItem)
	s.Step(`^I should receive an Item$`, is.iShouldReceiveAnItem)
}

