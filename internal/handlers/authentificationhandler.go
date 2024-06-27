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

	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	HashedPassword := []byte("pkenorngeongren")
	// Проверяем, соответствует ли введенный пароль хэшу
	err = bcrypt.CompareHashAndPassword(HashedPassword, []byte(user.Password))
	if err != nil {
		fmt.Println("Неправильный пароль:", err)
	}

	fmt.Println("Пароль верный! Вход разрешён")

	//secret jwt key generation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      user.Email,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("secret_key"))

	writer.Header().Set("Authorization", tokenString)

	response := "soooooo Good, bro"
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		print(err)
	}

	writer.WriteHeader(http.StatusOK)

}
