package initializers

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/tesauth?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")

	}

	log.Println("connecting to database")
}
