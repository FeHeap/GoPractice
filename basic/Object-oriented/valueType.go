package main

import "fmt"

type Human struct {
	name string
	age int8
	sex byte
}

func main() {
	//1. Initialize Human
	h1 := Human{ "Fe", 20, 1 }
	fmt.Printf("h1: %T , %v , %p \n", h1, h1, &h1)
	fmt.Println("-------------------")

	//2. Copy the struct object
	h2 := h1
	h2.name = "Zoe"
	h2.sex = 0
	fmt.Printf("After changing, h2: %T , %v , %p\n", h2, h2, &h2)
	fmt.Printf("h1: %T , %v , %p \n", h1, h1, &h1)
	fmt.Println("-------------------")

	//3. Transfer the struct object as parameter
	changeName(h1)
	fmt.Printf("h1: %T , %v , %p \n", h1, h1, &h1)
}

func changeName(h Human) {
	h.name = "Luder"
	h.age = 21
	fmt.Printf("In the function after changing, h: %T , %v , %p\n", h, h, &h)
}