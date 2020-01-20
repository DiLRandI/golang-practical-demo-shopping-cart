package integration_tests

import "github.com/DATA-DOG/godog"

type cartFeature struct {
}

func (c *cartFeature) iAmUsingCartService() error {
	return godog.ErrPending
}

func (c *cartFeature) iCallAddCartItem() error {
	return godog.ErrPending
}

func (c *cartFeature) iCallGetCartItems() error {
	return godog.ErrPending
}

func (c *cartFeature) iShouldReceiveResponseOfSuccess() error {
	return godog.ErrPending
}

func (c *cartFeature) iShouldReceivedTheAddedItemFromCart() error {
	return godog.ErrPending
}

func (c *cartFeature) iShouldReceiveEmptyCart() error {
	return godog.ErrPending
}

func (c *cartFeature) iCallClearCart() error {
	return godog.ErrPending
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
