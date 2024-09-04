package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/peter/go-auth/controllers"
)

func main() {
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/register", authcontroller.Register)

	fmt.Println("Server started at: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
