package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	// Exception 1
	res := math.Sqrt(-100)
	fmt.Println(res)
	res, err := Sqrt(-100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	// Exception 2
	res, err = Divide(100, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

	// Exception 3
	f, err := os.Open("/abc.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f.Name(), "該文件成功被打開！")
	}
	fmt.Println("====================")

	// 創建error對象的方法1
	err1 := errors.New("自己創建的錯誤！")
	fmt.Println(err1.Error())
	fmt.Println(err1)
	fmt.Printf("err1的類型: %T\n", err1) //*errors.errorString
	fmt.Println("--------------------")
	// 創建error對象的方法2
	err2 := fmt.Errorf("錯誤的類型%d", 10)
	fmt.Println(err2.Error())
	fmt.Println(err2)
	fmt.Printf("err2的類型: %T\n", err2) //*errors.errorString
	fmt.Println("--------------------")
	// error對象在函數中使用
	result, err3 := checkAge(-12)
	if err3 != nil {
		fmt.Println(err3.Error())
		fmt.Println(err3)
	} else {
		fmt.Println(result)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("負數不可以獲取平方根")
	} else {
		return math.Sqrt(f), nil
	}
}

func Divide(dividee float64, divider float64) (float64, error) {
	if divider == 0 {
		return 0, errors.New("出錯：除數不可為0！")
	} else {
		return dividee / divider, nil
	}
}

func checkAge(age int) (string, error) {
	if age < 0 {
		err := fmt.Errorf("您的年齡輸入是: %d, 該數字為負數, 有錯誤!", age)
		return "", err
	} else {
		return fmt.Sprintf("您的年齡輸入是: %d ", age), nil
	}
}