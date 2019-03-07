package function

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Result struct {
	Code    int
	Message string
	Results interface{}
}

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

const (
	SigningKey = "xxxxxxxxx"
)

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		panic(errors.New("Method not allowed"))
	}

	var userRegister UserRegister
	err := json.NewDecoder(r.Body).Decode(&userRegister)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Result{
			Code:    500,
			Message: "Parameters are required",
		})
	} else {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": userRegister.Email,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		})
		tokenString, err := token.SignedString([]byte(SigningKey))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Result{
				Code:    500,
				Message: err.Error(),
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(Result{
			Code:    200,
			Message: "OK",
			Results: struct{ token string }{
				token: tokenString,
			},
		})
		return
	}

}
