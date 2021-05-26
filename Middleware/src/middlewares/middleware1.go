package middlewares

import (
	"fmt"
	"net/http"
)

func Middleware1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middlewares 1:",r.URL)
		h.ServeHTTP(w, r)
	})
}
