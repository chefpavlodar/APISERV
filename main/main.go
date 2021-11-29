package main

import (
	"APISERV/console"
	"APISERV/logger"
	APISERV "APISERV/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port = ":50001"

//https://www.youtube.com/watch?v=z-mHhobE0Pw
func main() {

	logger.PrintAll()

	s := grpc.NewServer()
	srv := &console.GRPCServer{}
	APISERV.RegisterConsoleServer(s, srv)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
