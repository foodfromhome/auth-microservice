package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"nolabel-hac-auth-microservice-2024/internal/models"
	"time"
)

// Registration endpoint
func RegistrationUser(writer http.ResponseWriter, request *http.Request) {

	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	//TO DO: Проверка на то, что пользователя не существует в системе

	//password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Ошибка при генерации хэша пароля:", err)
	}

	user.Password = string(hashedPassword)
	user.Success = true

	//TO DO: add service methods

	//secret jwt key generation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      user.Email,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte("secret_key"))

	writer.Header().Set("Registration", tokenString)

	writer.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(writer).Encode(models.UserResponse{User: user})
	if err != nil {
		print(err)
	}
	writer.WriteHeader(http.StatusCreated)
}

func GetRegistration(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		// Отображение формы для ввода логина и пароля
		bytes, _ := ioutil.ReadFile("/Users/basty64/Programming/go/src/nolabel-hac-auth-microservice-2024/docs/page.html")
		fmt.Fprintf(writer, string(bytes))

	} else if request.Method == "POST" {
		// Обработка данных из формы
		username := request.FormValue("username")
		//password := request.FormValue("password")

		// Вывод сообщения об успешной аутентификации
		fmt.Fprintf(writer, "<h1>Hello, %s! You are now logged in.</h1>", username)
	}
}
