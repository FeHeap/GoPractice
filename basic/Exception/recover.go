package main

import "fmt"

func main() {
	fnA()
	fnB()
	fnC()
	fmt.Println("main over")
}

func fnA() {
	fmt.Println("This is fnA")
}

func fnB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("恢復啦，獲取recover的返回值:", msg)
		}
	}()
	fmt.Println("This is fnB")
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			panic("fnB 恐慌啦")
		}
	}
}

func fnC() {
	defer func() {
		fmt.Println("延遲執行函數")
		msg := recover()
		fmt.Println("獲取recover的返回值:", msg)
	}()
	fmt.Println("這是fnC")
	panic("fnC 恐慌啦")
}