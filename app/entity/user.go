package entity

import "github.com/golang-jwt/jwt/v4"

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	jwt.RegisteredClaims
}
