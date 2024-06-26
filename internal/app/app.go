package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nolabel-hac-auth-microservice-2024/internal/handlers"
)

type App struct {
	err    error
	router http.Handler
}

func NewApp() (*App, error) {
	a := &App{}

	return a, nil
}

func (a *App) HTTPRouter() http.Handler {

	router := mux.NewRouter()

	//Функция обработчик
	router.HandleFunc("/api/v1/registration", handlers.RegistrationUser).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/authentification", handlers.AuthentificationUser).Methods(http.MethodGet)

	a.router = router
	return a.router
}

func (a *App) Run() error {
	port := "localhost:8000"

	server := &http.Server{Addr: port, Handler: a.HTTPRouter()}

	log.Fatal(server.ListenAndServe())

	return nil
}
