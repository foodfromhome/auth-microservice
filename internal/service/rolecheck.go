package service

import "nolabel-hac-auth-microservice-2024/internal/models"

type RoleCheckService interface {
	CheckRole(role models.Role) (models.Role, error)
}

func CheckRole(role models.Role) models.Role {
	if role.Name == "admin" {
		return models.Role{
			Name:       "admin",
			Operations: []string{"create", "read", "update", "delete", "adminpanel"},
		}
	} else if role.Name == "cooker" {
		return models.Role{
			Name:       "user",
			Operations: []string{"create", "read"},
		}
	} else if role.Name == "delivery" {
		return models.Role{
			Name:       "delivery",
			Operations: []string{"read"},
		}
	}
	return models.Role{
		Name:       "user",
		Operations: []string{"read"},
	}
}
