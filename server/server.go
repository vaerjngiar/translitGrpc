package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"translitGrpc/translit"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("Error listening 8081 port", err)
	}

	server := grpc.NewServer()

	translit.RegisterTransliterationServer(server, NewTr())

	fmt.Println("Starting server at :8081")
	server.Serve(lis)
}
