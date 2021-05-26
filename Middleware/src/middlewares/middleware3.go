package middlewares
import (
	"net/http"
	"fmt"
)

func Middleware3_1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r*http.Request){
		fmt.Println("Middleware 3_1:",r.URL)
		h.ServeHTTP(w,r)
	})
}

func Middleware3_2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r*http.Request){
		fmt.Println("Middleware 3_2:",r.URL)
		h.ServeHTTP(w,r)
	})
}