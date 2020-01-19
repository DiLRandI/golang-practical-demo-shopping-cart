Feature: Item end point get
    In order to use Items I should be able
    to Get Items data

Scenario: When I call GetItems I should be able to see all the items saved .
	Given I am using Item Service
    When I call Get endpoint in Items
    Then I should receive Items