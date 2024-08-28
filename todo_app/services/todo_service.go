package services

import (
    "errors"
    "time"
    "github.com/dgrijalva/jwt-go/models"
    "github.com/dgrijalva/jwt-go/repository"
)



func CreateTodoList(todoList *models.TodoList) error {
    todoList.CreatedAt = time.Now()
    todoList.UpdatedAt = time.Now()
    return repositories.CreateTodoList(todoList)
}



func GetTodoLists(userId int) ([]models.TodoList, error) {
    return repositories.GetTodoLists(userId)
}



func UpdateTodoList(id int, userId int, todoList *models.TodoList) error {

    existing, err := repositories.GetTodoListById(id)
    if err != nil {
        return err
    }
    if existing.UserId != userId {
        return errors.New("unauthorized")
    }
    todoList.UpdatedAt = time.Now()
    return repositories.UpdateTodoList(id, todoList)
}



func DeleteTodoList(id int, userId int) error {

    existing, err := repositories.GetTodoListById(id)
    if err != nil {
        return err
    }
    if existing.UserId != userId {
        return errors.New("unauthorized")
    }
    now := time.Now()
    return repositories.DeleteTodoList(id, &now)
}



func CreateTodoMessage(listId int, userId int, todoMessage *models.TodoMessage) error {

    list, err := repositories.GetTodoListById(listId)
    if err != nil {
        return err
    }
    if list.UserId != userId {
        return errors.New("unauthorized")
    }
    todoMessage.CreatedAt = time.Now()
    todoMessage.UpdatedAt = time.Now()
    return repositories.CreateTodoMessage(listId, todoMessage)
}



func UpdateTodoMessage(listId int, messageId int, userId int, todoMessage *models.TodoMessage) error {

    list, err := repositories.GetTodoListById(listId)
    if err != nil {
        return err
    }
    if list.UserId != userId {
        return errors.New("unauthorized")
    }
    todoMessage.UpdatedAt = time.Now()
    return repositories.UpdateTodoMessage(messageId, todoMessage)
}



func DeleteTodoMessage(listId int, messageId int, userId int) error {

    list, err := repositories.GetTodoListById(listId)
    if err != nil {
        return err
    }
    if list.UserId != userId {
        return errors.New("unauthorized")
    }
    now := time.Now()
    return repositories.DeleteTodoMessage(messageId, &now)
}