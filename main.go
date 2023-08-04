package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jetsadawwts/go-gin-gorm/controllers"
	"github.com/jetsadawwts/go-gin-gorm/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.GET("/", controllers.Post)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.Get)
	r.GET("/posts/:id", controllers.GetById)
	r.PUT("/posts/:id", controllers.Put)
	r.DELETE("/posts/:id", controllers.Delete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
