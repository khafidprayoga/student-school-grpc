package main

import (
	"github.com/joho/godotenv"
	entity "github.com/khafidprayoga/grpc-basic/common/entities"
	"github.com/khafidprayoga/grpc-basic/common/utils"
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
		entity.PhoneNumber{},
	)
	if err != nil {
		panic(err)
	}
}
