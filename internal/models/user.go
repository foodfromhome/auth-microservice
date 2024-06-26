package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Login         string `json:"login"`
	Password      string `json:"password"`
	Success       bool   `json:"success"`
	StorageAccess string `json:"storage_access"`
}

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}
