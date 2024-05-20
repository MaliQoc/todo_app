package repositories

import (
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go/models"
)


var todoLists = []models.TodoList{}
var todoMessages = []models.TodoMessage{}


func CreateTodoList(todoList *models.TodoList) error {

    todoList.ID = len(todoLists) + 1
    todoLists = append(todoLists, *todoList)
    return nil
}



func GetTodoLists(userId int) ([]models.TodoList, error) {

    var userLists []models.TodoList
    for _, list := range todoLists {
        if list.UserId == userId && list.DeletedAt == nil {
            userLists = append(userLists, list)
        }
    }
    return userLists, nil
}



func GetTodoListById(id int) (*models.TodoList, error) {

    for _, list := range todoLists {
        if list.ID == id {
            return &list, nil
        }
    }
    return nil, errors.New("list not found")
}



func UpdateTodoList(id int, todoList *models.TodoList) error {

    for i, list := range todoLists {
        if list.ID == id {
            todoLists[i] = *todoList
            return nil
        }
    }
    return errors.New("list not found")
}



func DeleteTodoList(id int, deletedAt *time.Time) error {

    for i, list := range todoLists {
        if list.ID == id {
            todoLists[i].DeletedAt = deletedAt
            return nil
        }
    }
    return errors.New("list not found")
}



func CreateTodoMessage(listId int, todoMessage *models.TodoMessage) error {

    todoMessage.ID = len(todoMessages) + 1
    todoMessage.ListID = listId
    todoMessages = append(todoMessages, *todoMessage)
    return nil
}



func UpdateTodoMessage(id int, todoMessage *models.TodoMessage) error {

    for i, message := range todoMessages {
        if message.ID == id {
            todoMessages[i] = *todoMessage
            return nil
        }
    }
    return errors.New("message not found")
}



func DeleteTodoMessage(id int, deletedAt *time.Time) error {

    for i, message := range todoMessages {
        if message.ID == id {
            todoMessages[i].DeletedAt = deletedAt
            return nil
        }
    }
    return errors.New("message not found")
}