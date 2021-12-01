package console

import (
	APISERV "APISERV/proto"
	"context"
	"log"
)

//GRPCserver
type GRPCServer struct {
	APISERV.UnimplementedConsoleServer
}

//Add ....
func (s *GRPCServer) Add(ctx context.Context, req *APISERV.AddRequest) (*APISERV.AddResponse, error) {
	log.Println(req.GetX(), req.GetY())
	return &APISERV.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedConsoleServer() {}
