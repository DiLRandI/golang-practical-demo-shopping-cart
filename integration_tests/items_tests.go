package integration_tests

import (
	"fmt"
	"net/http"
	"github.com/DATA-DOG/godog"
)

func (a *apiFeature) iAmUsingItemService() error {
	a.itemServerUri="http://localhost:8080"
	return nil
}

func (a *apiFeature) iCallGetEndpointInItems() error {
	endpoint := "/getitem"
	var response interface{}
	rc,err := SendHttp(http.MethodGet,endpoint,a.itemServerUri,nil,&response,getServiceHeaders())
	a.err = err.Error()
	a.respCode = rc
	return nil
}

func (a *apiFeature) iShouldReceiveItemList() error {
	if a.respCode != 204 {
		return fmt.Errorf("Invalid Response")
	}
	return nil
}

func brandFeatureContext(s *godog.Suite, a *apiFeature) {
	s.Step(`^I am using Item Service$`, a.iAmUsingItemService)
	s.Step(`^I call Get endpoint in Items$`, a.iCallGetEndpointInItems)
	s.Step(`^I should receive Items$`, a.iShouldReceiveItemList)
}