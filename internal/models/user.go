package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	Success       bool   `json:"success"`
	StorageAccess string `json:"storage_access"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
