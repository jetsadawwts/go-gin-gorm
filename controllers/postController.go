package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jetsadawwts/go-gin-gorm/initializers"
	"github.com/jetsadawwts/go-gin-gorm/models"
)

func Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Start server post:" + " " + os.Getenv("PORT"),
	})
}

func PostsCreate(c *gin.Context) {

	// req body
	var p models.Post
	err := c.Bind(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// create post
	result := initializers.DB.Create(&p)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// return res
	c.JSON(201, gin.H{
		"message": p,
	})

}

func Get(c *gin.Context) {

	var p []models.Post

	initializers.DB.Find(&p)

	c.JSON(200, gin.H{
		"message": p,
	})

}

func GetById(c *gin.Context) {

	// req param
	id := c.Param("id")
	var p models.Post

	//get a post
	initializers.DB.First(&p, id)

	//return res
	c.JSON(200, gin.H{
		"message": p,
	})

}

func Put(c *gin.Context) {
	// req param
	id := c.Param("id")

	// req body
	var p models.Post
	c.Bind(&p)

	// get a post
	var get models.Post
	initializers.DB.First(&get, id)

	// update a post
	initializers.DB.Model(&get).Updates(&p)

	// return res
	c.JSON(200, gin.H{
		"message": p,
	})

}

func Delete(c *gin.Context) {
	// req param
	id := c.Param("id")


	var p models.Post

	// delete a post
	initializers.DB.Delete(&p ,id)

	// return res
	c.Status(200)

}
