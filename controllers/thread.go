package controllers

import (
    "net/http"
    "gossip-forum-backend/database"
    "gossip-forum-backend/models"
    "github.com/gin-gonic/gin"
)

func CreateThread(c *gin.Context) {
    var input models.Thread

		// Validates request by binding JSON payload to thread
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    thread := models.Thread{Title: input.Title, Content: input.Content, UserID: input.UserID}
    database.DB.Create(&thread)

    c.JSON(http.StatusOK, gin.H{"data": thread})
}

// Retrieves all threads from the database along with associated user information
func GetThreads(c *gin.Context) {
    var threads []models.Thread

    database.DB.Preload("User").Find(&threads)

    c.JSON(http.StatusOK, gin.H{"data": threads})
}

// Retrieves a single thread from the database along with associated user information, 
// throws an error if thread is not found
func GetThread(c *gin.Context) {
    var thread models.Thread
    if err := database.DB.Preload("User").Where("id = ?", c.Param("id")).First(&thread).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Thread not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": thread})
}

func UpdateThread(c *gin.Context) {
    var thread models.Thread
    if err := database.DB.Where("id = ?", c.Param("id")).First(&thread).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Thread not found"})
        return
    }

    var input models.Thread
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Model(&thread).Updates(input)

    c.JSON(http.StatusOK, gin.H{"data": thread})
}

func DeleteThread(c *gin.Context) {
	var thread models.Thread
	if err := database.DB.Where("id = ?", c.Param("id")).First(&thread).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Thread not found"})
			return
	}

	database.DB.Delete(&thread)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
