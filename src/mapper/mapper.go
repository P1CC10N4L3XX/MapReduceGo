package main

import (
	"MapReduceGo/src/config"
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
	"time"
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

func callReducer(chunk []int, address string, c chan string) {
	if chunk == nil {
		c <- fmt.Sprintf("Il chunk da inviare al reducer %s è nil...\n", address)
		return
	}
	req := &stubs.NumberChunk{Numbers: make([]int32, len(chunk))}
	for j, num := range chunk {
		req.Numbers[j] = int32(num)
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	utils.CheckError(err)
	defer conn.Close()

	client := stubs.NewReduceServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.ReduceChunk(ctx, req)
	c <- fmt.Sprintf("Reducer %s risponde: %s\n", address, resp.Status)
}

func splitChunk(numbers []int) ([]int, []int) {
	for i, n := range numbers {
		if n > config.REDUCER_SPLIT_NUMBER {
			return numbers[:i], numbers[i:]
		}
	}
	return numbers, nil
}

func (s *mapperServer) StartReduce(ctx context.Context, empty *emptypb.Empty) (*emptypb.Empty, error) {
	fmt.Println("Mapper: invia ai reducer il chunk ", numbers)
	chunk1, chunk2 := splitChunk(numbers)
	c := make(chan string)
	go callReducer(chunk1, config.REDUCER_ADDRESS[0], c)
	go callReducer(chunk2, config.REDUCER_ADDRESS[1], c)
	fmt.Println(<-c)
	fmt.Println(<-c)
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
