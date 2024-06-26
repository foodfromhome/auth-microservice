package handlers

import (
	"encoding/json"
	"net/http"
	"nolabel-hac-auth-microservice-2024/internal/models"
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
