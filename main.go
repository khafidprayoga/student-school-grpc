package main

import (
	"github.com/joho/godotenv"
	"github.com/khafidprayoga/grpc-basic/common/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	utils.GetDbConn()
}
