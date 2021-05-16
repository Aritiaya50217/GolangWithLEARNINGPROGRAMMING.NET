package main

import (
	"net/http"
	"test/controllers/account_controller"
)

func main() {

	http.HandleFunc("/account", account_controller.Index)
	http.HandleFunc("/account/index", account_controller.Index)

	http.HandleFunc("/account/login", account_controller.Login)
	http.HandleFunc("/account/welcome", account_controller.Welcome)
	http.HandleFunc("/account/logout", account_controller.Logout)
	//port
	http.ListenAndServe(":8080", nil)
}
