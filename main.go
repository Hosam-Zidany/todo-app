package main

import (
	"fmt"
	"todo-app/config"
	"todo-app/controllers"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("start")
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Todo{}, &models.User{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/todos", controllers.CreateTodo)
	r.GET("/todos", controllers.GetTodos)
	r.GET("/todo/:id", controllers.GetTodoById)
	r.PUT("/todo/:id", controllers.UpdateTodo)
	r.DELETE("todo/:id", controllers.DeleteTodo)

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.Run(":8080")
}
