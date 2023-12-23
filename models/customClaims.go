package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
