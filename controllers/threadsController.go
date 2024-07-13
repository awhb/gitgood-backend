package controllers

import (
	"net/http"

	"github.com/awhb/gitgood-backend/initialisers"
	"github.com/awhb/gitgood-backend/models"
	"github.com/gin-gonic/gin"
)

func ThreadsCreate(c *gin.Context) {
	// Get the authenticated user
	user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Current user does not exist"})
        return
    }
    authUser := user.(models.User)

    // Get data off request body
    var body models.Thread
    if c.ShouldBindJSON(&body) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Create a thread
    thread := models.Thread{
        Title:   body.Title,
        Content: body.Content,
        UserID:  authUser.ID, // Use the authenticated user's ID
		Tags:    body.Tags,
    }

	result := initialisers.DB.Create(&thread)

	// Check for errors
    if result.Error != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": result.Error.Error()})
        return
    }

    // Return created thread
    c.JSON(http.StatusOK, gin.H{"thread": thread})
}

// Retrieves all threads from the database
func ThreadsIndex(c *gin.Context) {
    var threads []models.Thread

    initialisers.DB.Preload("User").Preload("Comments").Find(&threads)

    c.JSON(http.StatusOK, gin.H{"threads": threads})
}

// Retrieves a single thread from the database along with associated user information
func ThreadsShow(c *gin.Context) {
    var thread models.Thread
    id := c.Param("id")

	// find thread
    initialisers.DB.Preload("User").Preload("Comments").First(&thread, id)

    if thread.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Thread not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"thread": thread})
}

func ThreadsUpdate(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

	var body models.Thread

    if c.ShouldBindJSON(&body) != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    // Find thread 
    var thread models.Thread
    initialisers.DB.First(&thread, id)

    // Update thread
    initialisers.DB.Model(&thread).Updates(body)

    c.JSON(http.StatusOK, gin.H{"thread": thread})
}

func ThreadsDelete(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Delete the thread
    initialisers.DB.Preload("User").Preload("Comments").Delete(&models.Thread{}, id)

    c.JSON(http.StatusOK, gin.H{"message": "Thread deleted successfully!"})
}
