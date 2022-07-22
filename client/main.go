package main

import (
	"context"
	"fmt"

	"github.com/khafidprayoga/grpc-basic/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.Dial("localhost:4500",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pb.NewStudentServiceClient(conn)

	listStudent, err := client.GetAllStudent(context.Background(), &emptypb.Empty{})
	fmt.Printf("List student:\n %v", listStudent)
}
