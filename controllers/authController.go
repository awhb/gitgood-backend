package controllers

import (
    // "github.com/awhb/gitgood-backend/models"
    // "github.com/awhb/gitgood-backend/utils"
    "github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// var input models.User

	// // Validates request by binding JSON payload to user
    // if err := c.ShouldBindJSON(&input); err != nil {
    //     c.JSON(, gin.H{"error": err.Error()})
    //     return
    // }

	// // Hashes user password
    // password, err := utils.HashPassword(input.Password)
    // if err != nil {
    //     c.JSON(, gin.H{"error": "Failed to hash password"})
    //     return
    // }

	// // Creates new user with provided username and hashed password
    // user := models.User{Username: input.Username, Password: password}
    // result := database.DB.Create(&user)

    // if result.Error != nil {
    //     c.JSON(, gin.H{"error": result.Error.Error()})
    //     return
    // }

    c.JSON(200, gin.H{"message": "Registration successful"})
}

func Login(c *gin.Context) {

	// Validates input by binding JSON payload
    // var input models.User
    // if err := c.ShouldBindJSON(&input); err != nil {
    //     c.JSON(, gin.H{"error": err.Error()})
    //     return
    // }

	// // Gets user from database
    // var user models.User
    // database.DB.Where("username = ?", input.Username).First(&user)

    // if user.ID == 0 {
    //     c.JSON(, gin.H{"error": "Invalid credentials"})
    //     return
    // }

	// // Checks if password is correct
    // if !utils.CheckPasswordHash(input.Password, user.Password) {
    //     c.JSON(, gin.H{"error": "Invalid credentials"})
    //     return
    // }

    c.JSON(200, gin.H{"message": "Login successful"})
}
