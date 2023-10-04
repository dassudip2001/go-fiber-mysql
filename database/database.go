package database

import (
	"fmt"
	"log"
	"os"

	"github.com/dassudip2001/webapp/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	godotenv.Load()
	// Get the database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	// // Construct the DSN string
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	// log.Println("Connecting to database...")
	// dsn := "admin:password@tcp(127.0.0.1:3306)/gofiber?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	} else {
		log.Println("Database connection successfully opened")
	}

	log.Println("Database connection successfully opened")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Run The Migrations")
	// add migrations here

	db.AutoMigrate(&models.Category{}, &models.Product{})
	// db.AutoMigrate(&models.Category{})

	Database = DbInstance{Db: db}
}
