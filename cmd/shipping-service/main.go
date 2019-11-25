package main

import (
	"context"
	"log"
	"net"

	shipping_service_pb "github.com/dilrandi/golang-practical-demo-shopping-cart/protos/shippingpb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting shipping service.")

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	log.Println("Start listing on port: ", 50051)
	s := grpc.NewServer()
	shipping_service_pb.RegisterShippingServiceServer(s, &grpcShipping{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type grpcShipping struct {
}

func (*grpcShipping) CalculateShippingCost(ctx context.Context, req *shipping_service_pb.CalculateShippingCostRequest) (*shipping_service_pb.CalculateShippingCostResponse, error) {
	log.Println("Calculating shipping cost for : ", req)
	//Shipping calculate logic.
	return nil, nil
}
