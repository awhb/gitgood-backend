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
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    authUser := user.(models.User)

    // Get data off request body
    var body struct {
        Title   string   `json:"title"`
        Content string   `json:"content"`
        Tags    []string `json:"tags"`
    }
    if err := c.BindJSON(&body); err != nil {
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

	result := initialisers.DB.Preload("User").Preload("Comments.User").Create(&thread)

	// Check for errors
    if result.Error != nil {
        c.Status(http.StatusForbidden)
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
    initialisers.DB.Preload("User").Preload("Comments.User").First(&thread, id)

    if thread.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Thread not found"})
        return
    }

	// load and return comments with thread
	var comments []models.Comment
		
	initialisers.DB.Preload("User").Preload("Thread").Find(&comments, "thread_id = ?", c.Param("id"))

    c.JSON(http.StatusOK, gin.H{"thread": thread, "comments": comments})
}

func ThreadsUpdate(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Get data off request body
    var body struct {
        Title string
        Content string
		Tags []string
    }

    c.ShouldBindJSON(&body)

    // Find thread 
    var thread models.Thread
    initialisers.DB.First(&thread, id)

    // Update thread
    initialisers.DB.Model(&thread).Updates(map[string]interface{}{
        "title": body.Title,
        "content": body.Content,
		"tags": body.Tags,
    })

    c.JSON(http.StatusOK, gin.H{"thread": thread})
}

func ThreadsDelete(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Delete the thread
    initialisers.DB.Delete(&models.Thread{}, id)

    c.JSON(http.StatusOK, gin.H{"message": "thread deleted"})
}
