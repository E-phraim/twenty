package handlers

import (
	"context"
	"net/http"
	"time"
	"twenty/db"
	"twenty/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var mongoClient *mongo.Client
var validate = validator.New()

func AllArticles(ctx *gin.Context) {
	cursor, err := db.MongoClient.Database("twenty-api").Collection("blogs").Find(context.TODO(), bson.D{{}})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var blogs []bson.M
	if err = cursor.All(context.TODO(), &blogs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

func OneArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var blog bson.M
	err = db.MongoClient.Database("twenty-api").Collection("blogs").FindOne(context.TODO(), bson.D{{"_id", articleID}}).Decode(&blog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

func CreateArticle(ctx *gin.Context) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var blog models.AddBlog
	defer cancel()

	if err := ctx.BindJSON(&blog); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if validationErr := validate.Struct(&blog); validationErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error validating all fields"})
		return
	}

	newBlog := models.Blog{
		Title:     blog.Title,
		Body:      blog.Body,
		Author:    blog.Author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	blogsCol := db.MongoClient.Database("twenty-api").Collection("blogs")

	result, err := blogsCol.InsertOne(context, newBlog)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)

}

func AggregateBlogs(c *gin.Context) {

	var pipeline interface{}
	if err := c.ShouldBindJSON(&pipeline); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	blogsCol := db.MongoClient.Database("twenty-api").Collection("blogs")

	cursor, err := blogsCol.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteArticle(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	id := c.Param("id")

	defer cancel()

	blogsCol := db.MongoClient.Database("twenty-api").Collection("blogs")

	articleID, _ := primitive.ObjectIDFromHex(id)

	result, err := blogsCol.DeleteOne(ctx, bson.M{"_id": articleID})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid article ID"})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"message": "article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article deleted"})

}

func UpdateArticle(c *gin.Context) {
	contxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blogsCol := db.MongoClient.Database("twenty-api").Collection("blogs")

	id := c.Param("id")
	articleID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid article ID"})
		return
	}

	var blog models.AddBlog
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding JSON"})
		return
	}

	if validatorErr := validate.Struct(&blog); validatorErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error validating input"})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"title":      blog.Title,
			"body":       blog.Body,
			"author":     blog.Author,
			"updated_at": time.Now(),
		},
	}

	result, err := blogsCol.UpdateOne(contxt, bson.M{"_id": articleID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating blog"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "blog not found"})
		return
	}

	var updatedBlog models.Blog
	err = blogsCol.FindOne(contxt, bson.M{"_id": articleID}).Decode(&updatedBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error retrieving updated blog"})
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}
