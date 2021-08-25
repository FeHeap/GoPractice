package main

import "fmt"

type Emp struct {
	name string
	age int8
	sex byte
}

func main() {

	emp1 := new(Emp)
	fmt.Printf("emp: %T , %v , %p \n", emp1, emp1, emp1)

	(*emp1).name = "Fe"
	(*emp1).age = 20
	(*emp1).sex = 1
	fmt.Println(emp1)

	emp1.name = "Luder"
	emp1.age = 21
	emp1.sex = 1
	fmt.Println(emp1)

	fmt.Println("---------------")

	SyntacticSugar()
}

func SyntacticSugar() {
	arr1 := [4]int{ 10, 20, 30, 40}
	arr2 := &arr1
	fmt.Println((*arr2)[len(arr1)-1])
	fmt.Println(arr2[0])
	arr3 := []int{ 100, 200, 300, 400 }
	arr4 := &arr3
	fmt.Println((*arr4)[len(arr3)-1])
}