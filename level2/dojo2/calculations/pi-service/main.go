package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	piProto "calculations/pb/pi/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port int

type server struct{}

var cache = map[uint64]uint64{}

func init() {
	flag.IntVar(&port, "port", 4000, "Port on which the RPC server is listening")
}

func main() {
	flag.Parse()

	host := fmt.Sprintf(":%d", port)
	conn, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	piProto.RegisterPiServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("Starting Pi Service server %s\n", host)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *piProto.Empty) (*piProto.PiResponse, error) {
	a := Pi()
	return &piProto.PiResponse{Result: a}, nil
}

func Pi() float32 {
	return 3.141592653589793238
}