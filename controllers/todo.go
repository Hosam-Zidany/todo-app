package controllers

import (
	"net/http"
	"todo-app/config"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var input models.Todo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := models.Todo{Title: input.Title, Completed: input.Completed}
	config.DB.Create(&todo)

	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	data := config.DB.Find(&todos)
	if data.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": data.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

