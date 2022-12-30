package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jwt-auth/config"
	"jwt-auth/models"
	"log"
	"strconv"
)

var (
	host     = config.GetDotEnvVar("DB_HOST")
	port, _  = strconv.Atoi(config.GetDotEnvVar("DB_PORT"))
	user     = config.GetDotEnvVar("DB_USER")
	password = config.GetDotEnvVar("DB_PASSWORD")
	dbname   = config.GetDotEnvVar("DB_NAME")
)

var DB *gorm.DB

func ConnectDB() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	fmt.Println("Successfully connected to database!")

	DB = conn
	conn.AutoMigrate(&models.User{})
}
