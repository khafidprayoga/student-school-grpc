package main

import (
	"context"
	"fmt"
	"time"

	"github.com/khafidprayoga/grpc-basic/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:4500",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pb.NewStudentServiceClient(conn)

	student := pb.CreateStudentRequest{
		FullName:  "Khafid Prayoga",
		Address:   "Mojokerto",
		Class:     "X-TKJ-2",
		BirthDate: uint64(time.Now().Unix()),
	}
	created, err := client.CreateStudent(context.Background(), &student)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	fmt.Printf("Data student\n\n%v \n", created)
}
