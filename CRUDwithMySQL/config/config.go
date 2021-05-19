package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {

	// ประเภท database
	dbDriver := "mysql"
	// ชื่อตาราง
	dbName := "db6"
	dbUser := "root"
	// รหัสผ่าน
	//dbPass := "123456"
	// รูปแบบการใช้ ถ้าหากมีการตั้งรหัสผ่าน
	//db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	db, err = sql.Open(dbDriver, dbUser+":"+"@/"+dbName)

	// connect to mysql
	// root:@/db6 ที่อยู่ของฐานข้อมูล
	//db, err = sql.Open("mysql", "root:@/db6")
	return

}
