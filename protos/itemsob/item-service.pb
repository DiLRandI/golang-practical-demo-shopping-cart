syntax = "proto3";
package services.items;
option go_package = "item-service.pb";

message GetItemByIDRequest {
   int32 ItemID=1;
}

message ItemResponse {
    float Cost = 1;
}

service ItemService {
    rpc GetItemByID (CalculateShippingCostRequest) returns (CalculateShippingCostResponse);
}