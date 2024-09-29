package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatalf("Failed to load environment file: %v", err)
	} else {
		fmt.Println("Environment file loaded successfully")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&Categories{}, &Images{}, &Products{}, &Roles{}, &Users{}, &Orders{}, &Cart{}, &OrderedList{})
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	categories := []Categories{
		{CategoryName: "Gadgets"},
		{CategoryName: "Furniture"},
		{CategoryName: "Make Up"},
		{CategoryName: "Tools"},
		{CategoryName: "Fashion"},
	}
	
	for _, category := range categories {
		var existingCategory Categories
		result := db.Where("category_name = ?", category.CategoryName).First(&existingCategory)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			log.Printf("Failed to check existing category: %v", result.Error)
			continue
		}
	
		if result.RowsAffected > 0 {
			fmt.Printf("Category %s already exists, skipping creation.\n", category.CategoryName)
		} else {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to create category %s: %v", category.CategoryName, err)
			} else {
				fmt.Printf("Category %s created successfully.\n", category.CategoryName)
			}
		}
	}
	

	DB = db
}
