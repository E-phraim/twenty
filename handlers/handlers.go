package handlers

import (
	"context"
	"net/http"
	"twenty/db"
	"twenty/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// var mongoClient *mongo.Client


func AllArticles(ctx *gin.Context){
	cursor, err := db.MongoClient.Database("twenty").Collection("blogs")
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var blogs []bson.M
	if err = cursor.All(context.TODO(), &blogs); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func OneArticle(ctx *gin.Context){
	id := ctx.Param("id")

	var blog bson.M
	err := db.MongoClient.Database("twenty").Collection("blogs")
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func CreateArticle(ctx *gin.Context){
	body := models.AddBlog{}

	if err := ctx.BindJSON(&body); err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var blog models.Blog

	blog.Title = body.Title
	blog.Body = body.Body
	blog.Author = body.Author

	if result := db.MongoClient.Database().
}

func DeleteArticle(ctx *gin.Context){

}

func UpdateArticle(ctx *gin.Context){

}

