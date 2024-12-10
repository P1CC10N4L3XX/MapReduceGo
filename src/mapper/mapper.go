package main

import (
	"MapReduceGo/src/protoBuffer/stubs"
	"MapReduceGo/src/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"strings"
)

type mapperServer struct {
	stubs.UnimplementedMapperServiceServer
}

func (s *mapperServer) ProcessChunk(ctx context.Context, chunk *stubs.NumberChunk) (*stubs.MapperResponse, error) {
	fmt.Printf("Mapper: ricevuto chunk: %s\n", chunk.Numbers)
	return &stubs.MapperResponse{Status: "Success"}, nil
}

func main() {
	address := os.Args[1]
	listener, err := net.Listen("tcp", strings.TrimPrefix(address, "localhost"))
	utils.CheckError(err)
	grpcServer := grpc.NewServer()
	stubs.RegisterMapperServiceServer(grpcServer, &mapperServer{})

	fmt.Println("Mapper server in ascolto su " + address)
	err = grpcServer.Serve(listener)
	utils.CheckError(err)

}
