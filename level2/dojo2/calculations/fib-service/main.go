package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	fibProto "calculations/pb/fib/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port int

type server struct{}

var cache = map[uint64]uint64{}

func init() {
	flag.IntVar(&port, "port", 3000, "Port on which the RPC server is listening")
}

func main() {
	flag.Parse()

	host := fmt.Sprintf(":%d", port)
	conn, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	fibProto.RegisterFibServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("Starting Fibonacci Service server %s\n", host)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *fibProto.FibRequest) (*fibProto.FibResponse, error) {
	log.Printf("fib-service: Compute method: a=%d\n", r.FibNum)
	a := Fibonacci(r.FibNum)

	return &fibProto.FibResponse{Result: a}, nil
}

func Fibonacci(n uint64) uint64 {
	if n <= 1 {
		return n
	  }
	  return Fibonacci(n-1) + Fibonacci(n-2)
}