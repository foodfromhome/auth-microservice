package models

type User struct {
	Email         string `json:"email"`
	Name          string `json:"name"`
	PasswordHash  string `json:"password_hash"`
	Success       bool   `json:"success"`
	StorageAccess string `json:"storage_access"`
}

type UserResponse struct {
	User User `json:"user"`
}
