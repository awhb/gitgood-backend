package controllers

import (
	"net/http"

	"github.com/awhb/gitgood-backend/initialisers"
	"github.com/awhb/gitgood-backend/models"
	"github.com/gin-gonic/gin"
)

func CommentsCreate(c *gin.Context) {
	// Get the authenticated user
	user, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }
    authUser := user.(models.User)

    // Get data off request body
    var body models.Comment
	
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

    // Create a thread
    comment := models.Comment{
        Content:   body.Content,
        ThreadID: body.ThreadID,
        UserID:  authUser.ID, // Use the authenticated user's ID
    }

	result := initialisers.DB.Create(&comment)

	// Check for errors
    if result.Error != nil {
        c.JSON(http.StatusForbidden, gin.H{"error": result.Error.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func CommentsUpdate(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Get data off request body
    var body models.Comment
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

    // Find the comment
    var comment models.Comment
    initialisers.DB.Preload("User").First(&comment, id)

    // Update the comment
    initialisers.DB.Model(&comment).Updates(body)

    c.JSON(http.StatusOK, gin.H{"comment": comment})
}

func CommentsDelete(c *gin.Context) {
    // Get the id off the url
    id := c.Param("id")

    // Delete the comment
    initialisers.DB.Preload("User").Delete(&models.Comment{}, id)

    c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}
