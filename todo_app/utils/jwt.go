package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your_secret_key")


type Claims struct {
    UserId int    `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}



func GenerateJWT(userId int, role string) (string, error) {

    claims := &Claims{
        UserId: userId,
        Role:   role,
        StandardClaims: jwt.StandardClaims{},
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}



func ParseJWT(tokenString string) (*Claims, error) {

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, errors.New("invalid token")
    }
    return claims, nil
}