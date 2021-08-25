package main

import "fmt"

type Teacher struct {
	name string
	age int8
	sex byte
}

func main() {
	var t1 Teacher
	fmt.Println(t1)
	fmt.Printf("t1: %T , %v , %q \n", t1, t1, t1)
	t1.name = "Fe"
	t1.age = 20
	t1.sex = 1
	fmt.Println(t1)
	fmt.Println("------------------")

	t2 := Teacher{}
	t2.name = "Luder"
	t2.age = 21
	t2.sex = 1
	fmt.Println(t2)
	fmt.Println("------------------")

	t3 := Teacher{
		name: "Sunny",
		age: 20,
		sex: 0,
	}
	t3 = Teacher{ name: "FeHeap", age: 30, sex: 1 }
	fmt.Println(t3)
	fmt.Println("------------------")

	t4 := Teacher{ "Rust", 11, 1 }
	fmt.Println(t4)
	fmt.Println("------------------")
}
