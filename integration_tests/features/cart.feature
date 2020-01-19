Feature: Cart End Points
    In order to use Items I should be able
    to Get Items data

Scenario: When I call add CartItem, I should be able to see my Item added in the cart .And when clear ,it should clear the Cart
	Given I am using Cart Service
    When I call add Cart Item 
    Then I should receive response of success
    When I call get CartItems
    Then I should received the added Item from cart
    When I call Clear Cart
    Then I should receive response of success
    When I call get CartItems
    Then I should receive empty cart