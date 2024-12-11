package main

import (
	"MapReduceGo/src/protoBuffer/stubs"
	"MapReduceGo/src/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"os"
	"sort"
	"strings"
)

type mapperServer struct {
	stubs.UnimplementedMapServiceServer
}

var numbers []int

func (s *mapperServer) MapChunk(ctx context.Context, chunk *stubs.NumberChunk) (*stubs.Reply, error) {
	for _, number := range chunk.Numbers {
		numbers = append(numbers, int(number))
	}
	fmt.Println("Mapper: esegue il sorting del chunk ", numbers, " ...")
	sort.Ints(numbers)
	fmt.Println("Chunk riordinato ", numbers)
	return &stubs.Reply{Status: "Success"}, nil
}

func (s *mapperServer) StartReduce(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	fmt.Println("Mapper: invia ai reducer il chunk ", numbers)
	return &emptypb.Empty{}, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: reducer <address:port>")
	}
	address := os.Args[1]
	listener, err := net.Listen("tcp", strings.TrimPrefix(address, "localhost"))
	utils.CheckError(err)
	grpcServer := grpc.NewServer()
	stubs.RegisterMapServiceServer(grpcServer, &mapperServer{})

	fmt.Println("Mapper server in ascolto su ", address)
	err = grpcServer.Serve(listener)
	utils.CheckError(err)

}
