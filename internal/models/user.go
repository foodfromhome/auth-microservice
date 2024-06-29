package models

type User struct {
	Name         string `json:"username"`
	PasswordHash string `json:"password"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}

type UserResponse struct {
	User User `json:"user"`
}

//type User struct {
//	Email         string `json:"email"`
//	Name          string `json:"name"`
//	PasswordHash  string `json:"password_hash"`
//	RoleName      Role   `json:"role_name"`
//	Success       bool   `json:"success"`
//	StorageAccess string `json:"storage_access"`
//}
