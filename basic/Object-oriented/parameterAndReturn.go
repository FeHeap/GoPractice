package main

import "fmt"

type Flower struct {
	name, color string
}

func main() {
	//1. parameter
	f1 := Flower{"玫瑰", "紅"}
	fmt.Printf("f1: %T , %v , %p \n", f1, f1, &f1)
	fmt.Println("-------------------")
	//struct as parameter
	changeInfo1(f1)
	fmt.Printf("f1: %T , %v , %p \n", f1, f1, &f1)
	fmt.Println("-------------------")
	//struct pointer as parameter
	changeInfo2(&f1)
	fmt.Printf("f1: %T , %v , %p \n", f1, f1, &f1)
	fmt.Println("-------------------")

	//2. return
	//return struct
	f2 := getFlower1()
	f3 := getFlower1()
	fmt.Println("Before changing,", f2, f3)
	fmt.Printf("The address of f2 is %p, and the adderss of f3 is %p.\n", &f2, &f3)
	f2.name = "杏花"
	fmt.Println("After changing,", f2, f3)
	fmt.Println("-------------------")
	//return struct pointer
	f4 := getFlower2()
	f5 := getFlower2()
	fmt.Println("Before changing,", f4, f5)
	f4.name = "桃花"
	fmt.Println("After changing,", f4, f5)
}

//transfer struct
func changeInfo1(f Flower) {
	f.name = "月季"
	f.color = "粉"
	fmt.Printf("In changeInfo1, f: %T , %v , %p \n", f, f, &f)
}

//transfer struct pointer
func changeInfo2(f *Flower) {
	f.name = "薔薇"
	f.color = "紫"
	fmt.Printf("In changeInfo2, f: %T , %v , %p , %p \n", f, f, f, &f)
}

func getFlower1() (f Flower) {
	f = Flower{"牡丹", "白"}
	fmt.Printf("In getFlower1, f: %T , %v , %p \n", f, f, &f)
	return
}

func getFlower2() (f *Flower) {
	//f = &Flower{"芙蓉", "紅色"}
	temp := Flower{"芙蓉", "紅色"}
	fmt.Printf("In getFlower2, temp: %T , %v , %p \n", temp, temp, &temp)
	f = &temp
	fmt.Printf("In getFlower2, f: %T , %v , %p , %p \n", f, f, f, &f)
	return
}