package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/yashnirmal/protoexample/proto"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuReq *pb.MenuRequest, srv grpc.ServerStreamingServer[pb.Menu]) error {
	items := []*pb.Item{
		{Id: "1", Name: "Espresso"},
		{Id: "2", Name: "Latte"},
	}

	for i,_ := range items {
		if err := srv.Send(&pb.Menu{Items: items[i : i+1]}); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) PlaceOrder(ctx context.Context,order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{OrderId: "1"}, nil
}

func (s *server) GetOrderStatus(ctx context.Context,receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: "1",
		Status:  "In Progress",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	
	if err != nil {
		log.Fatalf("Cannot create listener %s", err)
	}
	
	log.Printf("Server listening at %v", lis.Addr())
	
	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &server{})
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s", err)
	}
	
	fmt.Printf("Server started on port 8080\n")
}
