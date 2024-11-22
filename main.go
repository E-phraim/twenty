package main

import (
	"twenty/db"
	"twenty/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/api/v1/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hi I am Twenty; How do i help ?",
		})
	})

	db.ConnectMongoDB()

	routes.Routes(engine)

	engine.Run()

}
