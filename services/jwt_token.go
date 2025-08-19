package services

import (
 "fmt"
 "github.com/golang-jwt/jwt/v5"
 "time"
)

var secretKey = []byte("secret-key")

func createToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }
    fmt.Println("Token created successfully:", tokenString)
 return tokenString, nil
}