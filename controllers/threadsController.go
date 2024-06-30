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
		title string
		content string
		userId uint
		tags []string
    }
    c.Bind(&body)

    // Create a thread (and tags if they don't already exist)
	thread := models.Thread{
		Title:   body.title,
		Content: body.content,
		UserID:  body.userId,
	}

	result := initialisers.DB.Create(&thread)

	for _, tag := range body.tags {
		var t models.Tag
		initialisers.DB.Where("Name = ?", tag).FirstOrCreate(&t, models.Tag{Name: tag})
		thread.Tags = append(thread.Tags, t)
		initialisers.DB.Model(&thread).Association("Tags").Append(t)
	}

	// Check for errors
    if result.Error != nil {
        c.Status(http.StatusForbidden)
        return
    }

	threadResponse := models.MapThreadToResponse(thread)

    // Return created thread
    c.JSON(http.StatusOK, gin.H{"data": threadResponse})
}

// Retrieves all threads from the database along with associated user information
func ThreadsIndex(c *gin.Context) {
    var threads []models.Thread

    initialisers.DB.Preload("User").Preload("Tags").Find(&threads)

	threadResponse := models.MapThreadsToResponse(threads)

    c.JSON(http.StatusOK, gin.H{"data": threadResponse})
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

	// load and return comments
	var comments []models.Comment
		
	initialisers.DB.Preload("User").Preload("Thread").Find(&comments, "thread_id = ?", c.Param("id"))

	threadResponse := models.MapThreadToResponse(thread)
	commentsResponse := models.MapCommentsToResponse(comments)


    c.JSON(http.StatusOK, gin.H{"thread": threadResponse, "comments": commentsResponse})
}

func ThreadsUpdate(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Get data off request body
    var body struct {
        title string
        content string
    }
    c.Bind(&body)

    // Find thread 
    var thread models.Thread
    initialisers.DB.First(&thread, id)

    // Update thread
    initialisers.DB.Model(&thread).Updates(models.Thread{
        Title: body.title,
        Content: body.content,
    })

	threadResponse := models.MapThreadToResponse(thread)

    c.JSON(http.StatusOK, gin.H{"data": threadResponse})
}

func ThreadsDelete(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Delete the thread
    initialisers.DB.Delete(&models.Thread{}, id)

    c.JSON(http.StatusOK, gin.H{"message": "thread deleted"})
}

