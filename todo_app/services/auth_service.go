package services

import (
    "errors"
    "github.com/dgrijalva/jwt-go/models"
    "github.com/dgrijalva/jwt-go/utils"
)


type LoginCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}


var mockUsers = []models.User{
    {ID: 1, Username: "user1", Password: "password1", Role: "user"},
    {ID: 2, Username: "admin", Password: "password2", Role: "admin"},
}


func Authenticate(credentials LoginCredentials) (string, error) {
    for _, user := range mockUsers {
        if user.Username == credentials.Username && user.Password == credentials.Password {
            token, err := utils.GenerateJWT(user.ID, user.Role)
            if err != nil {
                return "", err
            }
            return token, nil
        }
    }
    return "", errors.New("invalid credentials")
}