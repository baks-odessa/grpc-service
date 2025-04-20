package main

import (
	"fmt"
	"grpc-service/internal/config"
)

func main() {
	// TODO: init loger
	// TODO: init app
	// TODO run grpc server

	cfg := config.MustLoad()

	fmt.Println(cfg)
}