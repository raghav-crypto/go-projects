package main

import (
	"gin/controllers"
	"gin/initializers"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	r := gin.Default()
	r.POST("/", controllers.PostsCreate)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.DELETE("/:id", controllers.DeletePost)
	r.PUT("/:id", controllers.UpdatePost)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"res": "pong",
		})
	})
	r.Run() // 3000
}
