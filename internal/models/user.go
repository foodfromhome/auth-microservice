package models

type User struct {
	Login         string `json:"login"`
	Password      string `json:"password"`
	Success       bool   `json:"success"`
	StorageAccess string `json:"storage_access"`
}
