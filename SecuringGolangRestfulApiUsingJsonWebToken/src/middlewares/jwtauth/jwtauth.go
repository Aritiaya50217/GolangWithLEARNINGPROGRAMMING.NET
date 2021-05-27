package jwtauth

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)
// สร้างตัวแปร 
var secretKey = "MySecretKey"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//  GET ค่า key
		tokenString := r.Header.Get("key")
		// ถ้า tokenString เป็นค่าว่าง
		if tokenString == "" {
			// เรียกใช้ฟังก์ชัน respondWithError 
			// แสดงข้อความ "Unauthorized"
			respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		} else {
			// ตัวแปร result จะเก็บข้อมูลที่ได้จากการเข้ารหัส Token เก็บข้อมูลแบบ interface ไม่จำกัดจำนวน
			result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// แสดงข้อมูลแบบ รหัสตัวเลขปนอักขระ เก็บข้อมูลแบบ []byte
			// secretKey ตัวแปรที่เก็บค่า key ของเรา
				return []byte(secretKey), nil
			})
			// ถ้าไม่มี error และ ตรวจสอบค่า result แล้ว
			if err == nil && result.Valid {
				// แสดงข้อมูล
				next.ServeHTTP(w,r)
			}else {
				respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			}
		}
	})
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
