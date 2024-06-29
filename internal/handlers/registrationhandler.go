package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
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
	url := fmt.Sprint("http://45.141.102.127:8090/api/users?email=", user.Email)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == http.StatusOK {
		writer.Header().Set("Registration", "User already registered")

		writer.Header().Set("Content-Type", "text/html; charset=utf-8")

		err = json.NewEncoder(writer).Encode("user registered")
		if err != nil {
			print(err)
		}
		writer.WriteHeader(http.StatusCreated)

	} else {

		//password hashing
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Ошибка при генерации хэша пароля:", err)
		}

		user.PasswordHash = string(hashedPassword)

		//TO DO: add service methods
		//user.RoleName = middleware.CheckRole(user.Role)

		//secret jwt key generation
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":      user.Email,
			"operations": user.Role,
			"expiration": time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, _ := token.SignedString([]byte("secret_key"))

		writer.Header().Set("Registration", tokenString)

		writer.Header().Set("Content-Type", "application/json")

		bytesRepresentation, err := json.Marshal(models.UserResponse{User: user})
		if err != nil {
			print(err)
		}

		writer.WriteHeader(http.StatusCreated)

		url := fmt.Sprint("http://45.141.102.127:8090/api/users?email=", user.Email)

		req, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			log.Println(err)
		}
		if req.StatusCode == http.StatusOK {
			writer.Header().Set("Status", "success")
		}

	}

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
