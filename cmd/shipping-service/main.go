package main

import (
	"context"
	"log"
	"net"

	shipping_service_pb "github.com/dilrandi/golang-practical-demo-shopping-cart/protos/shippingpb"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Llongfile)
	log.Println("Starting shipping service.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	shipping_service_pb.RegisterShippingServiceServer(s, &grpcShipping{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type grpcShipping struct {
}

func (*grpcShipping) CalculateShippingCost(ctx context.Context, req *shipping_service_pb.CalculateShippingCostRequest) (*shipping_service_pb.CalculateShippingCostResponse, error) {
	//Shipping calculate logic.
	return nil, nil
}
