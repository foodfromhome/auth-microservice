package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"nolabel-hac-auth-microservice-2024/internal/models"
	"nolabel-hac-auth-microservice-2024/internal/service"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Ошибка при генерации хэша пароля:", err)
	}

	user.PasswordHash = string(hashedPassword)
	user.Success = true

	//TO DO: add service methods
	user.RoleName = service.CheckRole(user.RoleName)

	//secret jwt key generation
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      user.Email,
		"operations": user.RoleName.Operations,
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

//1) По завершении регистрации отметьте учетную запись как неактивную (ожидает подтверждения) и создайте две строки из случайных символов.
//
//2) Сохраните обе строки в базе данных и свяжите их с пользователем.
//
//3) Отправьте пользователю ссылку по электронной почте, которая ведет на страницу вашего сайта и содержит обе строки.
//
//Пример: www.mysite.com/confirm.php?auth1=j0832r2&auth2=fji4j32ion
//
//4) Проверьте на своей странице, совпадают ли оба кода, и если да, пометьте учетную запись как активную.
