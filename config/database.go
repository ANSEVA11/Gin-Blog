package config

import(
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	dsn := "user=postgres password=1382 dbname=blogProject host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal("Database connection is faild ...!", err)
	}

	DB = db

	fmt.Println("Database connected successfully ...!")

}


