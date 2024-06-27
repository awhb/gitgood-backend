package controllers

import (
	"net/http"

	"github.com/awhb/gitgood-backend/initialisers"
	"github.com/awhb/gitgood-backend/models"
	"github.com/gin-gonic/gin"
)

func ThreadsCreate(c *gin.Context) {
    // Get data off request body
    var body struct {
		Title string
		Content string
		UserID uint
		tags []string
    }
    c.Bind(&body)

    // Create a thread (and tags if they don't already exist)
	thread := models.Thread{Title: body.Title, Content: body.Content, UserID: body.UserID}

	for _, tag := range body.tags {
		var t models.Tag
		initialisers.DB.Where("name = ?", tag).FirstOrCreate(&t, models.Tag{Name: tag})
		thread.Tags = append(thread.Tags, t)
	}

	result := initialisers.DB.Create(&thread)

	// add association between thread 

	// Check for errors
    if result.Error != nil {
        c.Status(http.StatusForbidden)
        return
    }

    // Return created thread
    c.JSON(http.StatusOK, gin.H{"data": thread})
}

// Retrieves all threads from the database along with associated user information
func ThreadsIndex(c *gin.Context) {
    var threads []models.Thread

    initialisers.DB.Preload("User").Preload("Tags").Find(&threads)

    c.JSON(http.StatusOK, gin.H{"data": threads})
}

// Retrieves a single thread from the database along with associated user information
func ThreadsShow(c *gin.Context) {
    var thread models.Thread
    id := c.Param("id")

    initialisers.DB.Preload("User").Preload("Comments.User").Preload("Tags").First(&thread, id)

    if thread.ID == 0 {
        c.Status(http.StatusNotFound)
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": thread})
}

func ThreadsUpdate(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Get data off request body
    var body struct {
        Title string
        Content string
    }
    c.Bind(&body)

    // Find thread 
    var thread models.Thread
    initialisers.DB.First(&thread, id)

    // Update thread
    initialisers.DB.Model(&thread).Updates(models.Thread{
        Title: body.Title,
        Content: body.Content,
    })

    c.JSON(http.StatusOK, gin.H{"data": thread})
}

func ThreadsDelete(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Delete the thread
    initialisers.DB.Delete(&models.Thread{}, id)

    c.JSON(http.StatusOK, gin.H{"data": "thread deleted"})
}
