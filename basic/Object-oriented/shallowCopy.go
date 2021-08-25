package main

import "fmt"

type Dog struct {
	name string
	color string
	age int8
	kind string
}

func main() {
	//1. Practice shallow copy of struct, default copy is shallow copy
	d1 := Dog{ "豆豆", "黑色", 2, "哈士奇" }
	fmt.Printf("d1: %T , %v , %p \n", d1, d1, &d1)
	d2 := d1 // Shallow Copy
	fmt.Printf("d2: %T , %v , %p \n", d2, d2, &d2)
	d2.name = "毛毛"
	fmt.Println("After changing, d2:", d2)
	fmt.Println("d1:", d1)
	fmt.Println("-------------------")

	//2. Practice shallow copy of struct, transferring point address directly
	d3 := &d1
	fmt.Printf("d3: %T , %v , %p \n", d3, d3, d3)
	d3.name = "球球"
	d3.color = "白色"
	d3.kind = "貴賓狗"
	fmt.Println("After changing, d3:", d3)
	fmt.Println("d1:", d1)
	fmt.Println("-------------------")

	//3. Practice shallow copy of struct through function new()
	d4 := new(Dog)
	d4.name = "多多"
	d4.color = "棕色"
	d4.age = 1
	d4.kind = "巴哥犬"
	d5 := d4
	fmt.Println("d4: %T , %v , %p \n", d4, d4, d4)
	fmt.Println("d5: %T , %v , %p \n", d5, d5, d5)
	fmt.Println("-------------------")
	d5.color = "金色"
	d5.kind = "黃金獵犬"
	fmt.Println("After changing, d5:", d5)
	fmt.Println("d4:", d4)
	fmt.Println("-------------------")
}