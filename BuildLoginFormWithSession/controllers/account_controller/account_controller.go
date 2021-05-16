package account_controller

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

// สร้างตัวแปร เก็บค่า sessions ที่สร้างขึ้นใหม่
var store = sessions.NewCookieStore([]byte("mysession"))

func Index(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("views/account_controller/index.html")
	tmp.Execute(w, nil)
}
func Login(w http.ResponseWriter, r *http.Request) {
	// รับค่าจาก form ในหน้า .html
	r.ParseForm()
	// ตัวแปรเก็บค่า
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	// ถ้า username ที่กรอกเข้ามาคือ abc และ password คือ 123 ให้แสดงผลหน้า index.html
	if username == "abc" && password == "123" {
		// การใช้ session ดึงค่าที่เราสร้างไว้ด้านบนมา
		session, _ := store.Get(r, "mysession")
		// ใช้ session เก็บค่า value ที่ชื่อ username ที่เราประกาศไว้ด้านบน
		session.Values["username"] = username
		//การ save  session
		session.Save(r, w)
		http.Redirect(w, r, "/account/welcome", http.StatusSeeOther)

	} else {
		// สร้างการ map[]
		data := map[string]interface{}{
			"err": "Invalid",
		}
		tmp, _ := template.ParseFiles("views/account_controller/index.html")
		tmp.Execute(w, data)
	}
}
func Welcome(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")
	username := session.Values["username"]

	data := map[string]interface{}{
		"username": username,
	}
	tmp, _ := template.ParseFiles("views/account_controller/welcome.html")
	tmp.Execute(w, data)
}
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")
	// อายุของ session
	session.Options.MaxAge = -1
	session.Save(r, w)
	// เมื่อ logout แล้วจะไปยังหน้า index
	http.Redirect(w, r, "/account/index", http.StatusSeeOther)
}
