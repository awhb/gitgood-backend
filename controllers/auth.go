package controllers

import (
    "net/http"
    "gossip-forum-backend/database"
    "gossip-forum-backend/models"
    "gossip-forum-backend/utils"
    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

		// Validates input by binding JSON payload
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

		// Hashes user password
    password, err := utils.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

		// Creates new user with provided username and hashed password
    user := models.User{Username: input.Username, Password: password}
    result := database.DB.Create(&user)

    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {

		// Validates input by binding JSON payload
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

		// Gets user from database
    var user models.User
    database.DB.Where("username = ?", input.Username).First(&user)

    if user.ID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

		// Checks if password is correct
    if !utils.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
