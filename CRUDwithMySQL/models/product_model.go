package models

import (
	"test/config"
	"test/entities"
)

type ProductModel struct{}

// สร้าง func FindAll
func (*ProductModel) FindAll() ([]entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		// product คือ ชื่อตารางที่เราสร้างในฐานข้อมูล mysql
		rows, err2 := db.Query("select * from product")
		if err2 != nil {
			return nil, err2
		} else {
			var products []entities.Product
			for rows.Next() {
				var product entities.Product
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
				products = append(products, product)
			}
			return products, nil
		}
	}
}

func (*ProductModel) Create(product *entities.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		// product คือ ชื่อตารางที่เราสร้างในฐานข้อมูล mysql
		// values(?,?,?,?) ดึงค่ามาใสใน database
		result2, err2 := db.Exec("insert in product(name,price,quantity,description) values(?,?,?,?)", product.Name, product.Price, product.Quantity, product.Description)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result2.RowsAffected()
			return rowsAffected > 0
		}
	}
}

// การลบข้อมูลแต่ละ id
func (*ProductModel) Delete(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		// product คือ ชื่อตารางที่เราสร้างในฐานข้อมูล mysql
		result, err2 := db.Exec("delete from product where id = ?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}

// ค้นหาตาม id
func (*ProductModel) Find(id int64) (entities.Product, error) {
	db, err := config.GetDB()
	if err != nil {
		// return ข้อมูลที่เราสร้างใน Product struct
		return entities.Product{}, err
	} else {
		// แสดงข้อมูลทั้งหมดของตาราง product โดยต้องใส่ id
		rows, err2 := db.Query("select * from product where id=?", id)
		if err2 != nil {
			return entities.Product{}, err2
		} else {
			var product entities.Product
			for rows.Next() {
				rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Description)
			}
			return product, nil
		}
	}
}

// การอัพเดทข้อมูล

func (*ProductModel) Update(product entities.Product) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		// where id=? คือ กำหนด id ที่ต้องการอัพเดท
		result, err2 := db.Exec("Update product set name =?,price=?,quantity=?,description=?,where id=?", product.Name, product.Price, product.Quantity, product.Description, product.Id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}
