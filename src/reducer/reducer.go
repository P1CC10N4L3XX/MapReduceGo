package main

import (
	"MapReduceGo/src/config"
	"MapReduceGo/src/protoBuffer/stubs"
	"MapReduceGo/src/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var path string

type reducerServer struct {
	stubs.UnimplementedReduceServiceServer
	mutex sync.Mutex
}

func fileToSlice(data string) []int {
	content := strings.TrimSpace(data)
	if content == "" {
		return nil
	}
	numStrs := strings.Split(content, ",")

	var numbers []int

	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		utils.CheckError(err)
		numbers = append(numbers, num)
	}

	return numbers
}

func (s *reducerServer) ReduceChunk(ctx context.Context, numberChunk *stubs.NumberChunk) (*stubs.Reply, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var chunk []int
	for _, number := range numberChunk.Numbers {
		chunk = append(chunk, int(number))
	}
	fmt.Println("Reducer: riceve il chunk ", chunk, " ...")
	fileRead, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0666)
	utils.CheckError(err)
	defer fileRead.Close()

	data := make([]byte, 1024)
	n, err := fileRead.Read(data)
	if err != io.EOF {
		utils.CheckError(err)
	}

	numbers := fileToSlice(string(data[:n]))
	for _, number := range chunk {
		numbers = append(numbers, number)
	}
	sort.Ints(numbers)
	fmt.Println("Reducer: scrive sul file ", numbers, "...")
	var strNumbers []string
	for _, number := range numbers {
		strNumbers = append(strNumbers, fmt.Sprintf("%d", number))
	}
	output := strings.Join(strNumbers, ",")
	fileWrite, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	utils.CheckError(err)
	defer fileWrite.Close()
	_, err = fileWrite.WriteString(output)
	utils.CheckError(err)
	return &stubs.Reply{Status: "Success"}, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: reducer <address:port>")
	}
	address := os.Args[1]
	if address == config.REDUCER_ADDRESS[0] {
		path = "../output/reducer1.txt"
	} else if address == config.REDUCER_ADDRESS[1] {
		path = "../output/reducer2.txt"
	}
	listener, err := net.Listen("tcp", strings.TrimPrefix(address, "localhost"))
	utils.CheckError(err)
	grpcServer := grpc.NewServer()
	stubs.RegisterReduceServiceServer(grpcServer, &reducerServer{})

	fmt.Println("Reducer server in ascolto su ", address)
	err = grpcServer.Serve(listener)
	utils.CheckError(err)
}
