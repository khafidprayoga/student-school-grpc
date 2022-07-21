package utils

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDbConn() *gorm.DB {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
