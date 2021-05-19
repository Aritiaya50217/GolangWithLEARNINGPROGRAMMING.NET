package authen_middleware

import (
	"net/http"
	"github.com/gorilla/sessions"
)
// สร้างตัวแปรเก็บ ค่า session ใหม่
var store = sessions.NewCookieStore([]byte("mysession"))

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter,r*http.Request){
		session,_ := store.Get(r,"mysession")
		username := session.Values["username"]
		if username == nil {
			// ถ้า username มีค่าว่าง หรือ ไม่มีการ login ให้ไปยังหน้า /login 
			http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
		}else {
			h.ServeHTTP(w,r)
		}
	})
}