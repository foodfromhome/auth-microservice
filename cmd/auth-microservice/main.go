package main

import (
	"fmt"
	"log"
	"nolabel-hac-auth-microservice-2024/internal/app"
)

func main() {
	fmt.Println("hola amigos")

	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
