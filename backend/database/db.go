package database

import (
	"fmt"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() error {
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	println("DB_USERNAME:", DB_USERNAME)
	println("DB_PASSWORD:", DB_PASSWORD)
	println("DB_NAME:", DB_NAME)
	println("DB_HOST:", DB_HOST)
	println("DB_PORT:", DB_PORT)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	println("dsn:", dsn)
	//fmt.Println("dsn:", dsn)

	println("Connecting to database. . . . . .")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connecting to database. error = %w", err)
		//os.Exit(1)
	}

	println("Connected to database!")

	DB = db
	return nil

}
