package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/peter/go-auth/entities"
	"github.com/peter/go-auth/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
}

var userModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tem.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/auth/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
		var user entities.User
		userModel.Where(&user, "username", UserInput.Username)

		var message error
		if user.Username == "" {
			message = errors.New("username tidak ditemukan")
		} else {
			// membandingan user input dengan di database
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Username), []byte(UserInput.Password))
			if errPassword != nil {
				message = errors.New("password salah")
			}
		}

		if message != nil {
			temp, _ := template.ParseFiles("views/auth/login.html")
			temp.Execute(w, message)
		} else {
			temp, _ := template.ParseFiles("views/auth/login.html")
			temp.Execute(w, nil)
		}

		if message == nil {
			data := map[string]interface{}{
				"error": message,
			}
			temp, _ := template.ParseFiles("views/auth/login.html")
			temp.Execute(w, data)
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("views/auth/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tem.Execute(w, nil)
}
