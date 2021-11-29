package console

import (
	APISERV "APISERV/proto"
	"context"
)

//GRPCserver
type GRPCServer struct {
	APISERV.UnimplementedConsoleServer
}

//Add ....
func (s *GRPCServer) Add(ctx context.Context, req *APISERV.AddRequest) (*APISERV.AddResponse, error) {
	return &APISERV.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedConsoleServer() {}
