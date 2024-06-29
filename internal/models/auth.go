package models

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleName Role   `json:"role_name"`
}

type AuthResponse struct {
	Success bool `json:"success"`
}
