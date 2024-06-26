package controllers

import (
    "net/http"
    "gossip-backend/models"
    "github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
    var input models.Comment
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    comment := models.Comment{Content: input.Content, ThreadID: input.ThreadID, UserID: input.UserID}
    database.DB.Create(&comment)

    c.JSON(http.StatusOK, gin.H{"data": comment})
}

func GetComments(c *gin.Context) {
    var comments []models.Comment
    database.DB.Preload("User").Where("thread_id = ?", c.Param("thread_id")).Find(&comments)

    c.JSON(http.StatusOK, gin.H{"data": comments})
}

func UpdateComment(c *gin.Context) {
    var comment models.Comment
    if err := database.DB.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
        return
    }

    var input models.Comment
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Model(&comment).Updates(input)

    c.JSON(http.StatusOK, gin.H{"data": comment})
}

func DeleteComment(c *gin.Context) {
    var comment models.Comment
    if err := database.DB.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
        return
    }

    database.DB.Delete(&comment)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
