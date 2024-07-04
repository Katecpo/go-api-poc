package api

import (
	"awesomeProject/api/models"
	"awesomeProject/api/models/requests"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("GET /users")

	users := models.Users{
		models.User{Email: "test@test.com", Username: "bdespierres"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		errFormated := fmt.Errorf("error while serializing response")
		fmt.Println(errFormated.Error())
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /user")

	var requestBody requests.CreateUserRequest
	if decodeBody(&requestBody, w, r) != nil {
		return
	}

	user := models.User{Email: requestBody.Email, Username: requestBody.Username}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		errFormated := fmt.Errorf("error while serializing response")
		fmt.Println(errFormated.Error())
	}
}

func decodeBody[T any](body *T, w http.ResponseWriter, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return err
	}

	return nil
}
