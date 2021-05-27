package middlewares

import (
	"fmt"
	"net/http"
)

// การสร้าง admin โดยใช้ middlewares
func BasicAuth(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// สร้างตัวแปร เก็บ username password และ ok
		username, password, ok := r.BasicAuth()
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
		fmt.Println("ok: ", ok)
		// || คือ or
		if !ok || !CheckUsernameAndPassword(username, password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Account Invalid"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized\n"))
			return
		}
		handle(w, r)
	}
}

func CheckUsernameAndPassword(username, password string) bool {
	return username == "admin" && password == "12345"
}
