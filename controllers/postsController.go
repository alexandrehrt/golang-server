package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title, Body string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if (result.Error != nil) {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off the body
	var body struct {
		Title, Body string
	}

	c.Bind(&body)

	// Find the post to update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}

