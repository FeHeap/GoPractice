package main

import (
	"fmt"
	"time"
)

func main() {
	time1 := time.Now()
	TestTime()
	time2 := time.Now()
	fmt.Println(time2.Sub(time1).Seconds())
}

func TestTime() {
	t := time.Now()
	fmt.Println("1.", t)
	fmt.Println("2.", t.Local())
	fmt.Println("3.", t.UTC())
	t = time.Date(2018, time.January, 1, 1, 1,1, 0, time.Local)
	fmt.Printf("4. 本地時間: %s, 國際統一時間: %s\n", t, t.UTC())
	t, _ = time.Parse("2006-01-02 15:04:05", "2018-07-19 05:47:13")
	fmt.Println("5.", t)
	fmt.Println("6.", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("7.", time.Now().String())
	fmt.Println("8.", time.Now().Unix())
	fmt.Println("9.", time.Now().UnixNano())
	fmt.Println("10.", t.Equal(time.Now()))
	fmt.Println("11.", t.Before(time.Now()))
	fmt.Println("12.", t.After(time.Now()))
	year, month, day := time.Now().Date()
	fmt.Printf("13. %T %v, %T %v, %T %v\n", year, year, month, month, day, day)
	fmt.Println("14.", time.Now().Year())
	fmt.Println("15.", time.Now().Month())
	fmt.Println("16.", time.Now().Day())
	fmt.Println("17.", time.Now().Weekday())
	hour, minute, second := time.Now().Clock()
	fmt.Printf("18. %T %v, %T %v, %T %v\n", hour, hour, minute, minute, second, second)
	fmt.Println("19.", time.Now().Hour())
	fmt.Println("20.", time.Now().Minute())
	fmt.Println("21.", time.Now().Second())
	fmt.Println("22.", time.Now().Nanosecond())
	fmt.Println("23.", time.Now().Sub(time.Now()))
	fmt.Println("24.", time.Now().Sub(time.Now()).Hours())
	fmt.Println("25.", time.Now().Sub(time.Now()).Minutes())
	fmt.Println("26.", time.Now().Sub(time.Now()).Seconds())
	fmt.Println("27.", time.Now().Sub(time.Now()).Nanoseconds())
	fmt.Println("28.", "時間間距:", t.Sub(time.Now()).String())
	d, _ := time.ParseDuration("1h30m")
	fmt.Println("29.", d)
	fmt.Println("30.", "交卷時間:", time.Now().Add(d))
	fmt.Println("31.", "一年一個月零一天之後的日期:", time.Now().AddDate(1, 1, 1))
}