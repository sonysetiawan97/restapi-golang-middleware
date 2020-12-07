package controllers

import (
	"net/http"

	"example.com/v1/models"
	"github.com/gin-gonic/gin"
)

// GetTodos ...
func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "todo not found",
			"status":  http.StatusNotFound,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"results": todo,
			"message": "todo founded",
			"status":  http.StatusOK,
		})
	}
}

// CreateATodo ...
func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "todo not found",
			"status":  http.StatusNotFound,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"results": todo,
			"message": "todo founded",
			"status":  http.StatusOK,
		})
	}
}

// GetATodo ...
func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo models.Todo
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "todo not found",
			"status":  http.StatusNotFound,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"results": todo,
			"message": "todo founded",
			"status":  http.StatusOK,
		})
	}
}

// UpdateATodo ...
func UpdateATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "todo not found",
			"status":  http.StatusNotFound,
		})
	}
	c.BindJSON(&todo)
	err = models.UpdateATodo(&todo, id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "todo not found",
			"status":  http.StatusNotFound,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"results": todo,
			"message": "todo founded",
			"status":  http.StatusOK,
		})
	}
}

// DeleteATodo ...
func DeleteATodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	err := models.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "todo not found",
			"status":  http.StatusNotFound,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"results": "id " + id + " is deleted",
			"message": "todo founded",
			"status":  http.StatusOK,
		})
	}
}
