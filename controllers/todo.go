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

func GetTodoById(c *gin.Context) {

	var todo models.Todo
	id := c.Param("id")

	result := config.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	result := config.DB.First(&todo, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.Title = input.Title
	todo.Completed = input.Completed
	config.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	result := config.DB.First(&todo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}

	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
