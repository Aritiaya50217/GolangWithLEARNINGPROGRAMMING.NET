package main

import (
	"fmt"
	"time"
)

func Demo1() {
	today := time.Now()
	fmt.Println("Today:", today)

	year := today.Year()
	month := today.Month()
	day := today.Day()
	hour := today.Hour()
	minites := today.Minute()
	second := today.Second()

	fmt.Println("year:", year)
	fmt.Println("month:", month)
	fmt.Println("Index of month:", int(month))
	fmt.Println("day:", day)
	fmt.Println("hour:", hour)
	fmt.Println("minites:", minites)
	fmt.Println("second:", second)
}
func Demo2() {
	today := time.Now()
	fmt.Println("Format date:", today.Format("01/02/2006"))
	fmt.Println("Format time:", today.Format("15:04:05"))
	fmt.Println("Format date and time:", today.Format("01/02/2006 15:04:05"))
}
func Demo3() {
	s := "10/20/1980"
	value, err := time.Parse("01/02/2006", s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("New date:", value.Format("2006-01-02"))
	}
}
func Demo4() {
	date1 := time.Date(2019, time.Month(01), 03, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2019, time.Month(01), 05, 0, 0, 0, 0, time.UTC)
	days := date2.Sub(date1).Hours() / 24
	fmt.Print("Days:", days)

}

// แสดงวันล่วงหน้า
func Demo5() {
	// วันเวลาปัจจุบัน
	today := time.Now()
	// เพิ่มไป 2 วัน จากวันปัจจุบัน
	// AddDate(0, 0, 2) --> (year,month,day)
	newDate := today.AddDate(0, 0, 2)
	fmt.Println("Today:", today)
	fmt.Println("New date:", newDate)
	// แบบที่ 2
	newDate2 := today.Add(2 * 24 * time.Hour)
	fmt.Println("Nwe date 2:", newDate2)
}

func main() {
	// Demo1()
	// Demo2()
	// Demo3()
	// Demo4()
	Demo5()
}
