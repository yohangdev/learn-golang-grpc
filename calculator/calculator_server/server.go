package main

import (
	"context"
	"fmt"
	"github.com/yohang88/learn-golang-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}

func (s server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v", request)

	firstInteger := request.GetFirstInteger()
	secondInteger := request.GetSecondInteger()

	result := firstInteger + secondInteger

	response := &calculatorpb.SumResponse{
		SumResult: result,
	}

	return response, nil
}

func main()  {
	fmt.Printf("CalculatorService start to listening...")

	listener, err := net.Listen("tcp", "0.0.0.0:50002")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(rpcServer, &server{})

	err = rpcServer.Serve(listener)

	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}