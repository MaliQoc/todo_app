package controllers

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go/utils"
)

func AuthMiddleware(c *gin.Context) {

    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
        c.Abort()
        return
    }

    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    if tokenString == authHeader {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
        c.Abort()
        return
    }

    claims, err := utils.ParseJWT(tokenString)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        c.Abort()
        return
    }

    c.Set("userId", claims.UserId)
    c.Set("role", claims.Role)
    c.Next()
}