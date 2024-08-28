package main

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go/controllers"
)

func main() {
    r := gin.Default()

    r.POST("/login", controllers.Login)

    authorized := r.Group("/api")
    authorized.Use(controllers.AuthMiddleware)
    {
        authorized.POST("/todo", controllers.CreateTodoList)
        authorized.GET("/todos", controllers.GetTodoLists)
        authorized.PUT("/todo/:id", controllers.UpdateTodoList)
        authorized.DELETE("/todo/:id", controllers.DeleteTodoList)
        authorized.POST("/todo/:id/message", controllers.CreateTodoMessage)
        authorized.PUT("/todo/:id/message/:messageId", controllers.UpdateTodoMessage)
        authorized.DELETE("/todo/:id/message/:messageId", controllers.DeleteTodoMessage)
    }

    r.Run(":8080")
}