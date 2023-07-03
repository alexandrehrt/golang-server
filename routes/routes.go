package routes

import (
	"go-crud/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostShow)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)
}