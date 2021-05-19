package main

import (
	"net/http"
	"test/controllers/admin/category_controller"
	"test/controllers/admin/login_controller"
	"test/controllers/admin/product_controller"
	"test/controllers/home_controller"
	"test/controllers/middlewares/authen_middleware"
)

func main() {
	http.HandleFunc("/", home_controller.Index)
	http.HandleFunc("/home", home_controller.Index)
	http.HandleFunc("/home/index", home_controller.Index)

	// ใช้ func AuthMiddleware เพื่อบังคับให้มีการ login
	http.HandleFunc("/admin/category", authen_middleware.AuthMiddleware(category_controller.Index))
	http.HandleFunc("/admin/category/index", authen_middleware.AuthMiddleware(category_controller.Index))

	http.HandleFunc("/admin/product", authen_middleware.AuthMiddleware(product_controller.Product))
	http.HandleFunc("/admin/product/index", authen_middleware.AuthMiddleware(product_controller.Product))

	// login
	http.HandleFunc("/admin/login", login_controller.UserLogin)
	http.HandleFunc("/admin/login/index", login_controller.UserLogin)
	http.HandleFunc("/admin/login/process", login_controller.Process)

	// logout
	http.HandleFunc("/admin/logout", login_controller.UserLogout)

	// port
	http.ListenAndServe(":8080", nil)
}
