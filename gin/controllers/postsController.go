package controllers

import (
	"fmt"
	"gin/initializers"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostsCreate(c *gin.Context) {
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts := models.Post{Title: input.Title, Body: input.Body}
	initializers.DB.Create(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetPost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	var post models.Post
	if result := initializers.DB.First(&post, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
	}
	fmt.Println(post)
	c.JSON(http.StatusOK, &post)
}

func DeletePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	var post models.Post

	if result := initializers.DB.First(&post, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	initializers.DB.Delete(&post)
	c.Status(http.StatusOK)
}
func UpdatePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	body := models.Post{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var post models.Post

	if result := initializers.DB.First(&post, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Save(&post)
	c.JSON(http.StatusOK, &post)
}
