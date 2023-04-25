package main

import (
	"encoding/json"
	"fmt"
	"gin/controllers"
	"gin/initializers"
	"net/http"

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
	r.GET("/ping", func(c *gin.Context) {
		url := "http://api.open-notify.org/astros.json"
		res, err := http.Get(url)
		if err != nil {
			fmt.Println("error getting response", err)
			return
		}
		defer res.Body.Close()
		var response AstronautResponse
		err = json.NewDecoder(res.Body).Decode(&response)

		if err != nil {
			fmt.Println("error reading response body", err)
			http.Error(c.Writer, "Error decoding response body", http.StatusInternalServerError)

			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"res": json.NewEncoder(c.Writer).Encode(response),
		})
	})
	r.Run()
}
