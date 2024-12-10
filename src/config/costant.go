package config

var MAPPER_ADDRESS = []string{
	"localhost:50051",
	"localhost:50052",
	"localhost:50053",
	"localhost:50054",
}

var REDUCER_ADDRESS = []string{
	"localhost:50055",
	"localhost:50056",
}

var MAPPER_NUMBER = len(MAPPER_ADDRESS)
var REDUCER_NUMBER = len(REDUCER_ADDRESS)
