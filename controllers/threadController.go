package controllers

import (
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
        
    }
    c.Bind(&body)

    // Create a thread
    thread := models.Thread{Title: body.Title, Content: body.Content, UserID: body.UserID}
    
    result := initialisers.DB.Create(&thread)

    if result.Error != nil {
        c.Status(400)
        return
    }

    // Return created thread
    c.JSON(200, gin.H{"data": thread})
}

// Retrieves all threads from the database along with associated user information
func ThreadsIndex(c *gin.Context) {
    var threads []models.Thread

    initialisers.DB.Preload("User").Find(&threads)

    c.JSON(200, gin.H{"data": threads})
}

// Retrieves a single thread from the database along with associated user information
func ThreadsShow(c *gin.Context) {
    var thread models.Thread
    id := c.Param("id")

    initialisers.DB.Preload("User").First(&thread, id)

    if thread.ID == 0 {
        c.Status(404)
        return
    }

    c.JSON(200, gin.H{"data": thread})
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

    c.JSON(200, gin.H{"data": thread})
}

func ThreadsDelete(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Delete the thread
    initialisers.DB.Delete(&models.Thread{}, id)

    c.Status(200)
}
