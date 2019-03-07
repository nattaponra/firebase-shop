package function

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Error struct {
	Error string
}
type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		panic(errors.New("Method not allowed"))
	}

	var userRegister UserRegister
	err := json.NewDecoder(r.Body).Decode(&userRegister)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Error{
			Error: "Parameters are required",
		})
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(userRegister)
	}

}
