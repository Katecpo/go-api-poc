package main

import (
	"awesomeProject/api"
	"fmt"
	"net/http"
)

func main() {
	startApi()
}

func listen() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", api.GetUsers)
	mux.HandleFunc("POST /user", api.CreateUser)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}

func startApi() {
	listen()
}
