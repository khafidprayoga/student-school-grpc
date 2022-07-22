package server

import "google.golang.org/grpc"

func GetGrpcServer() *grpc.Server {
	return grpc.NewServer()
}
