package main

import (
	"fmt"
	"time"
)

//1. 定義結構體，表示自定義錯誤類型
type MyError struct {
	When time.Time
	What string
}

//2. 實現Error()方法
func (e MyError) Error() string {
	return fmt.Sprintf("%v : %v", e.When, e.What)
}

//3. 定義函數，返回error對象。該函數求矩形面積
func getArea(width, length float64) (float64, error) {
	errorInfo := ""
	if width < 0 && length < 0 {
		errorInfo = fmt.Sprintf("長度: %v, 寬度: %v, 均為負數", length, width)
	} else if length < 0 {
		errorInfo = fmt.Sprintf("長度: %v, 出現負數", length)
	} else if width < 0 {
		errorInfo = fmt.Sprintf("寬度: %v, 出現負數", width)
	}

	if errorInfo != "" {
		return 0, MyError{time.Now(), errorInfo}
	} else {
		return width * length, nil
	}
}

func main() {
	area, err := getArea(-4, -5)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Println("面積:", area)
	}
}