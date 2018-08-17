package models 

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.db

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DBUSER")
	password := os.Getenv("PASS")
	dbName := os.Getenv("DBNAME")
	dbHost := os.Getenv("DBHOST")

	bUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Sneaker{})
}

func GetDB() *gorm.DB {
	return db
}