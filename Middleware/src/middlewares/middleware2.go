package middlewares

import (
	"fmt"
	"net/http"
)

func Middleware2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middlewares 2:", r.URL)
		h.ServeHTTP(w, r)
	})
}
