package main

import (
	"github.com/jetsadawwts/go-gin-gorm/initializers"
	"github.com/jetsadawwts/go-gin-gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}


