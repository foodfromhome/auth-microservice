package testdata

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
