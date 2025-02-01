package handlers

import (
	"context"
	"gin-api/helpers"
	"gin-api/models"
	"gin-api/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate status
	err := helpers.ValidateTaskStatus(task.Status)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		task.ID = primitive.NewObjectID()

		_, err = utils.DB.Database(os.Getenv("DB_NAME")).Collection("tasks").InsertOne(context.Background(), task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
	}

}

func GetTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := utils.DB.Database(os.Getenv("DB_NAME")).Collection("tasks").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing tasks"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTaskByID fetches a user by their ID
func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	err = utils.DB.Database(os.Getenv("DB_NAME")).Collection("tasks").FindOne(context.Background(), bson.M{"id": taskID}).Decode(&task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	userID := c.Param("id")

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedTask bson.M

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"id": objID}
	update := bson.M{"$set": updatedTask}

	_, err = utils.DB.Database(os.Getenv("DB_NAME")).Collection("tasks").UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	taskID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	_, err = utils.DB.Database(os.Getenv("DB_NAME")).Collection("tasks").DeleteOne(context.Background(), bson.M{"id": taskID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// GetPaginatedTasks retrieves tasks with pagination
func GetPaginatedTasks(c *gin.Context) {
	ctx := context.Background()
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	skipInt64 := int64((page - 1) * limit)

	// Convert to int64
	limitInt64 := int64(limit)

	// Create options with correct int64 types
	findOptions := options.Find().SetSkip(skipInt64).SetLimit(limitInt64)

	cursor, err := utils.DB.Database(os.Getenv("DB_NAME")).Collection("tasks").Find(
		ctx,
		bson.M{},
		findOptions,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode tasks"})
		return
	}

	// Get total count for pagination
	total, err := utils.DB.Database(os.Getenv("DB_NAME")).Collection("tasks").CountDocuments(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count total documents"})
		return
	}

	totalPages := (int(total) + limit - 1) / limit

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"pagination": gin.H{
			"current_page": page,
			"total_pages":  totalPages,
			"total_items":  total,
			"limit":        limit,
		},
	})
}
