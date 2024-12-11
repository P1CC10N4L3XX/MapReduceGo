package main

import (
	"MapReduceGo/src/protoBuffer/stubs"
	"MapReduceGo/src/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strings"
)

var numbers []int

type reducerServer struct {
	stubs.UnimplementedReduceServiceServer
}

func (s *reducerServer) ReduceChunk(ctx context.Context, chunk *stubs.NumberChunk) (*stubs.Reply, error) {

	fmt.Println("Reducer esegue il merge sorting del chunk ", numbers, " ...")
	return &stubs.Reply{Status: "Success"}, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: reducer <address:port>")
	}
	address := os.Args[1]
	listener, err := net.Listen("tcp", strings.TrimPrefix(address, "localhost"))
	utils.CheckError(err)
	grpcServer := grpc.NewServer()
	stubs.RegisterReduceServiceServer(grpcServer, &reducerServer{})

	fmt.Println("Reducer server in ascolto su ", address)
	err = grpcServer.Serve(listener)
	utils.CheckError(err)
}
