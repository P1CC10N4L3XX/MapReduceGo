package main

import (
	"MapReduceGo/src/config"
	"MapReduceGo/src/protoBuffer/stubs"
	"MapReduceGo/src/utils"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func fileToSlice(data string) []int {
	content := strings.TrimSpace(data)

	numStrs := strings.Split(content, ",")

	var numbers []int

	for _, numStr := range numStrs {
		num, err := strconv.Atoi(numStr)
		utils.CheckError(err)
		numbers = append(numbers, num)
	}

	return numbers
}

func chunkSlice(numbers []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		chunks = append(chunks, numbers[i:end])
	}
	return chunks
}

func callMapper(address string, chunk []int, c chan string) {
	req := &stubs.NumberChunk{Numbers: make([]int32, len(chunk))}
	for j, num := range chunk {
		req.Numbers[j] = int32(num)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	utils.CheckError(err)
	defer conn.Close()

	client := stubs.NewMapperServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.ProcessChunk(ctx, req)
	utils.CheckError(err)
	c <- fmt.Sprintf("Mapper %s risponde: %s\n", address, resp.Status)

}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: master <file_name>")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	utils.CheckError(err)
	defer file.Close()

	data := make([]byte, 1024)
	n, err := file.Read(data)
	utils.CheckError(err)

	numbers := fileToSlice(string(data[:n]))
	chunks := chunkSlice(numbers, config.MAPPER_NUMBER)
	c := make(chan string)
	for i, chunk := range chunks {
		go callMapper(config.MAPPER_ADDRESS[i], chunk, c)
	}
	for i := 0; i < len(chunks); i++ {
		fmt.Println(<-c)
	}
}
