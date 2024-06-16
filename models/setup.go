package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("Failed to load environment file")
	} else {
		fmt.Println("Environment file loaded successfully")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
										os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&Categories{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&Images{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&Products{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&Roles{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&Users{})
	if err != nil {
		return
	}

	DB = db
}