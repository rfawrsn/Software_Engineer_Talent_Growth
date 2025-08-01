package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"task-api/config"
	"task-api/models"
)

var taskCollection *mongo.Collection

func InitTaskController() {
	taskCollection = config.DB.Collection("tasks")
}

func CreateTask(c *gin.Context) {
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	if input.Priority != "Low" && input.Priority != "Medium" && input.Priority != "High" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Priority must be Low, Medium, or High"})
		return
	}

	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	res, err := taskCollection.InsertOne(context.TODO(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created", "id": res.InsertedID})
}

func GetTasks(c *gin.Context) {
	filter := bson.M{}

	if cat := c.Query("category"); cat != "" {
		filter["category"] = cat
	}
	if prio := c.Query("priority"); prio != "" {
		filter["priority"] = prio
	}

	sortParam := c.Query("sort")
	sortOrder := c.DefaultQuery("order", "asc")

	sort := bson.D{}
	if sortParam != "" {
		order := 1
		if sortOrder == "desc" {
			order = -1
		}
		sort = append(sort, bson.E{Key: sortParam, Value: order})
	}

	opts := options.Find().SetSort(sort)
	cursor, err := taskCollection.Find(context.TODO(), filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(context.TODO())

	tasks := make([]models.Task, 0)
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var task models.Task
	err = taskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	input.UpdatedAt = time.Now()
	update := bson.M{"$set": input}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated models.Task
	err = taskCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": objID}, update, opts).Decode(&updated)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Update failed or task not found"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	res, err := taskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil || res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or already deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task successfully deleted"})
}
