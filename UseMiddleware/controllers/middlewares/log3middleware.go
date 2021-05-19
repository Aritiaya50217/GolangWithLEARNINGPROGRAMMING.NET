package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Log3Middleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		today := time.Now()
		fmt.Println("Date: " + today.Format("01/02/2006 15:04:05"))
		h.ServeHTTP(w, r)
	})
}
