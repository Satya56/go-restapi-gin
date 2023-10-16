package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Open new database connection
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal(".env file couldn't be loaded")
	}
	DBHOST := os.Getenv("DBHOST")
	DBPORT := os.Getenv("DBPORT")
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	DBNAME := os.Getenv("DBNAME")

	connstring := DBUSER + ":" + DBPASS + "@(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=UTC"
	database, err := gorm.Open(mysql.Open(connstring))
	fmt.Println(connstring)

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
