package database

import (
    "log"
    "web-forum-backend/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=postgres dbname=web_forum port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    database.AutoMigrate(&models.User{}, &models.Thread{}, &models.Comment{}, &models.Tag{})

    DB = database
}
