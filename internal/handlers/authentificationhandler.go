package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"nolabel-hac-auth-microservice-2024/internal/models"
	"time"
)

// Authentification endpoint
func AuthentificationUser(writer http.ResponseWriter, request *http.Request) {

	var auth models.Auth

	err := json.NewDecoder(request.Body).Decode(&auth)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	//TO DO: create request
	HashedPassword := []byte("requesttouserservice")

	// Проверяем, соответствует ли введенный пароль хэшу
	err = bcrypt.CompareHashAndPassword(HashedPassword, []byte(auth.Password))
	if err != nil {
		fmt.Println("Неправильный пароль")
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}
	fmt.Println("Пароль верный! Вход разрешён")

	//secret jwt key generation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      auth.Email,
		"operations": auth.RoleName.Operations,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("secret_key"))

	writer.Header().Set("Authorization", tokenString)

	writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(models.AuthResponse{Success: true})
	if err != nil {
		print(err)
	}

	writer.WriteHeader(http.StatusOK)

}
