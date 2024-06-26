package controllers

import (
    "github.com/awhb/gossip-backend/initialisers"
    "github.com/awhb/gossip-backend/models"
    "github.com/gin-gonic/gin"
)

func CommentsCreate(c *gin.Context) {
    // Get data off request body
    var body struct {
        Content string
        UserID uint
        ThreadID uint
    }
    c.Bind(&body)

    // Create a comment
    comment := models.Comment{Content: body.Content, UserID: body.UserID, ThreadID: body.ThreadID}
    
    result := initialisers.DB.Create(&comment)

    if result.Error != nil {
        c.Status(400)
        return
    }
}

func CommentsIndex(c *gin.Context) {
    var comments []models.Comment
    
    initialisers.DB.Preload("User").Preload("Thread").Find(&comments, "thread_id = ?", c.Param("thread_id"))

    c.JSON(200, gin.H{"comments": comments})
}

func CommentsUpdate(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Get data off request body
    var body struct {
        Content string
    }
    c.Bind(&body)

    // Find the comment
    var comment models.Comment
    initialisers.DB.First(&comment, id)

    // Update the comment
    initialisers.DB.Model(&comment).Updates(models.Comment{
        Content: body.Content,
    })
    
    c.JSON(200, gin.H{"comment": comment})
}

func CommentsDelete(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Delete the comment
    initialisers.DB.Delete(&models.Comment{}, id)

    c.Status(200)
}
