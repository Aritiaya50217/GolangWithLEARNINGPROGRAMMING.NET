package accountapi

import (
	"fmt"
	"net/http"
	"test/entities"
	"time"

	"encoding/json"

	jwt "github.com/dgrijalva/jwt-go"
)

// สร้างตัวแปร
var secretKey = "MySecretKey"

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	// ดึงข้อมูลจาก Account struct
	var account entities.Account
	// ใช้ package json
	// r.body คือ ส่วนของ body ใน postman
	// Decode(&account) อ่านข้อมูลแบบ json ที่รับค่ามาจาก &account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			// from entities.Account
			"username": account.Username,
			"password": account.Password,
			// create time
			"exp": time.Now().Add(time.Hour * 72).Unix(),
		})
		// token.SignedString(key interface{}) (string, error)
		tokenString, err2 := token.SignedString([]byte(secretKey))
		if err2 != nil {
			// call func respondWithError
			respondWithError(w, http.StatusBadRequest, err2.Error())
		} else {
			// call func respondWithJson
			respondWithJson(w, http.StatusOK, entities.JWTToken{Token: tokenString})
		}
	}
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("key")
	// ตัวแปร result จะเก็บข้อมูลที่ได้จากการเข้ารหัส Token เก็บข้อมูลแบบ interface ไม่จำกัดจำนวน
	//  token ตัวแปรที่สร้างไว้ด้านบน
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// แสดงข้อมูลแบบ รหัสตัวเลขปนอักขระ เก็บข้อมูลแบบ []byte
		// secretKey ตัวแปรที่เก็บค่า key ของเรา
		return []byte(secretKey), nil
	})
	// ถ้าตรวจสอบแล้ว ค่าก่อนและหลัง เข้ารหัส Token เท่ากัน แสดงคำว่า Valid
	if err == nil && result.Valid {
		fmt.Println("Valid")
	} else {
		// ค่าก่อนและหลัง เข้ารหัส Token ไม่เท่ากัน แสดงคำว่า Invalid
		fmt.Println("InValid")
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	// เรียกใช้งานฟังก์ชัน respondWithJson
	// สร้าง amp เก็บข้อความ key : value 
	// "error" :"Unauthorized" โดย "Unauthorized" มาจาก func JWTAuth ใน jwtauth.go
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	// payload เป็นตัวแปรที่เก็บข้อมูลจาก entities.JWTToken{Token: tokenString}
	response, _ := json.Marshal(payload)
	// .Set(key string, value string)
	w.Header().Set("Content-Type", "application/json")
	// code เป็นตัวแปรที่แสดงข้อมูล http.StatusOK เช่น Status: 200 OK
	// 200 เป็น int
	w.WriteHeader(code)
	
	w.Write(response)
}
