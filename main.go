package main

import (
	"fmt"
	"net"

	"github.com/joho/godotenv"
	entity "github.com/khafidprayoga/grpc-basic/common/entities"
	"github.com/khafidprayoga/grpc-basic/common/utils"
	"github.com/khafidprayoga/grpc-basic/controller"
	"github.com/khafidprayoga/grpc-basic/proto/pb"
	"github.com/khafidprayoga/grpc-basic/server"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := utils.GetDbConn()
	db.AutoMigrate(
		entity.Student{},
	)
	if err != nil {
		panic(err)
	}

	srv := controller.NewStudentServer(db)
	s := server.GetGrpcServer()
	pb.RegisterStudentServiceServer(s, &srv)

	lis, err := net.Listen("tcp", ":4500")
	if err != nil {
		panic(err)
	}
	fmt.Println("GRPC Server running")
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
