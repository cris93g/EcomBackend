package db

import (
	"fmt"
	"log"
	"os"

	"github.com/cris93g/backendEcom/pkg/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)


var (
	DBConn *gorm.DB
)


func InitDatabase(){
	e:=	godotenv.Load()
  	if e != nil {
    	log.Fatal("Error loading .env file")
  	}
  	connectionUrl := os.Getenv("CONNECTION_URL")
	var err error
	DBConn, err = gorm.Open("postgres", connectionUrl)
	if err != nil {
		panic("failed to connect to db")
	}
	fmt.Println("database is connected")

	DBConn.AutoMigrate(&models.GoProducts{})
	fmt.Println("database migrated")
}