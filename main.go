package main

import (
	"github.com/gin-gonic/gin"
	"./app"
)

func Router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		// /books
		v1.GET("/books", app.GetBooks)
		v1.POST("/books", app.CreateBook)

		// /books/:id
		v1.GET("/book/:id", app.GetBook)
		v1.PUT("/book/:id", app.UpdateBook)
		v1.DELETE("/book/:id", app.DeleteBook)
	}

	return router
}

func main() {
	app.InitDB()
	router := Router()
	router.Run(":5000")
}
