package login_controller

import (
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

// สร้างตัวแปรเก็บค่า sessions
var store = sessions.NewCookieStore([]byte("mysession"))

func UserLogin(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("views/admin/login/login.html")
	tmp.Execute(w, nil)
}
func Process(w http.ResponseWriter, r *http.Request) {
	// ดึงค่าที่กรอกอยู่ใน form
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	if username == "abc" && password == "123" {
		session, _ := store.Get(r, "mysession")
		// เก็บ session จาก username
		session.Values["username"] = username
		// บันทึกค่า session ที่กรอกมา
		session.Save(r, w)
		// หลังจาก login แล้วให้ไปยังหน้า /admin/category
		http.Redirect(w, r, "/admin/category", http.StatusSeeOther)

	} else {
		// สร้าง map เก็บข้อความ
		data := map[string]interface{}{
			// key : value
			"err": "Invalid",
		}
		tmp, _ := template.ParseFiles("views/admin/login/login.html")
		tmp.Execute(w, data)
	}

}

func UserLogout(w http.ResponseWriter,r*http.Request){
	session,_ := store.Get(r,"mysession")
	session.Options.MaxAge = -1
	session.Save(r,w)
	// เมื่อ logout ไปยังหน้า login
	http.Redirect(w,r,"/admin/login",http.StatusSeeOther)
}