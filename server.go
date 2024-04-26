package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "grpc_server/proto"
)

var (
	port = flag.Int("port", 8888, "The server port")
)

type server struct {
	pb.UnimplementedGetMaxServer
}

func (s *server) FromTwo(ctx context.Context, req *pb.GetMaxRequest) (*pb.GetMaxResponse, error) {
	max := math.Max(float64(req.Num1), float64(req.Num2))
	intMax := int32(max)

	return &pb.GetMaxResponse{
		MaxNum: intMax,
	}, nil

}

func (s *server) FromList(ctx context.Context, req *pb.GetListMaxRequest) (*pb.GetMaxResponse, error) {
	max := int32(0)
	for _, num := range req.Nums {
		if num > max {
			max = num
		}
	}
	return &pb.GetMaxResponse{
		MaxNum: max,
	}, nil

}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetMaxServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
