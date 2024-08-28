package controllers

import (
    "net/http"
    "github.com/dgrijalva/jwt-go/models"
    "github.com/dgrijalva/jwt-go/services"
    "github.com/gin-gonic/gin"
    "strconv"
)



func CreateTodoList(c *gin.Context) {

    var todoList models.TodoList
    if err := c.ShouldBindJSON(&todoList); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    userId := c.MustGet("userId").(int)
    todoList.UserId = userId
    
    if err := services.CreateTodoList(&todoList); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todoList)
}



func GetTodoLists(c *gin.Context) {

    userId := c.MustGet("userId").(int)
    todoLists, err := services.GetTodoLists(userId)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todoLists)
}



func UpdateTodoList(c *gin.Context) {

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var todoList models.TodoList
    if err := c.ShouldBindJSON(&todoList); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userId := c.MustGet("userId").(int)
    if err := services.UpdateTodoList(id, userId, &todoList); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todoList)
}



func DeleteTodoList(c *gin.Context) {

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    userId := c.MustGet("userId").(int)
    if err := services.DeleteTodoList(id, userId); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo list deleted"})
}



func CreateTodoMessage(c *gin.Context) {

    listId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var todoMessage models.TodoMessage
    if err := c.ShouldBindJSON(&todoMessage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userId := c.MustGet("userId").(int)
    if err := services.CreateTodoMessage(listId, userId, &todoMessage); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todoMessage)
}



func UpdateTodoMessage(c *gin.Context) {

    listId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    messageId, err := strconv.Atoi(c.Param("messageId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Message ID"})
        return
    }

    var todoMessage models.TodoMessage
    if err := c.ShouldBindJSON(&todoMessage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userId := c.MustGet("userId").(int)
    if err := services.UpdateTodoMessage(listId, messageId, userId, &todoMessage); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todoMessage)
}



func DeleteTodoMessage(c *gin.Context) {

    listId, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    messageId, err := strconv.Atoi(c.Param("messageId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Message ID"})
        return
    }

    userId := c.MustGet("userId").(int)
    if err := services.DeleteTodoMessage(listId, messageId, userId); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Todo message deleted"})
}