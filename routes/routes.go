package routes

import (
	"twenty/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/api/v1")

	routes.GET("/timeline", handlers.AllArticles)
	routes.GET("/article/:id", handlers.OneArticle)

	routes.POST("/article", handlers.CreateArticle)
	routes.DELETE("/article/:id", handlers.DeleteArticle)
	routes.PUT("/article/:id", handlers.UpdateArticle)

	routes.POST("/blogs/aggregations", handlers.AggregateBlogs)
}
