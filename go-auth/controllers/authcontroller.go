package controllers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tem.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("views/auth/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tem.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	tem, err := template.ParseFiles("views/auth/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tem.Execute(w, nil)
}