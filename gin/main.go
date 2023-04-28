package main

import (
	"gin/controllers"
	"gin/initializers"

	"github.com/gin-gonic/gin"
)

type Astronaut struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}
type AstronautResponse struct {
	Message string      `json:"message"`
	Number  int         `json:"number"`
	People  []Astronaut `json:"people"`
}

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	r := gin.Default()
	r.POST("/", controllers.PostsCreate)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.DELETE("/:id", controllers.DeletePost)
	r.PUT("/:id", controllers.UpdatePost)

	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.Login)
	r.Run()
}
