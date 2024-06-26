package handlers

import (
	"encoding/json"
	"net/http"
	"nolabel-hac-auth-microservice-2024/internal/models"
	"text/template"
)

func RegistrationUser(writer http.ResponseWriter, request *http.Request) {

	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	//service methods
	response := "Good, bro"
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		print(err)
	}

	writer.WriteHeader(http.StatusCreated)
}

func AuthentificationUser(writer http.ResponseWriter, request *http.Request) {

	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	//service methods
	response := "soooooo Good, bro"
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		print(err)
	}

	writer.WriteHeader(http.StatusOK)

}

var (
	tmpl = template.Must(template.ParseFiles("/Users/basty64/Programming/go/src/nolabel-hac-auth-microservice-2024/docs/page.html"))
)

func Testing(w http.ResponseWriter, r *http.Request) {
	data := models.User{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	//data.Success = true
	//data.StorageAccess = "Hello, bro!"
	tmpl.Execute(w, data)
}
