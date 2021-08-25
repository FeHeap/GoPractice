package main

import "fmt"

type A interface {
}

type Cat struct {
	name string
	age int
}

type Hito struct {
	name string
	sex string
}

func main() {
	var a1 A = Cat{"Mimi", 1}
	var a2 A = Hito{"Steven", "male"}
	var a3 A = "Learn golang with me!"
	var a4 A = 100
	var a5 A = 3.14
	showInfo(a1)
	showInfo(a2)
	showInfo(a3)
	showInfo(a4)
	showInfo(a5)
	fmt.Println("-------------------")

	fmt.Println("The parameter of println is empty interface. It can be any type.", 100, 3.14, Cat{"Neko", 2})

	map1 := make(map[string]interface{})
	map1["name"] = "Daniel"
	map1["age"] = 13
	map1["height"] = 1.71
	fmt.Println("-------------------")

	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, a5)
	fmt.Println(slice1)
}

func showInfo(a A) {
	fmt.Printf("%T , %v\n", a, a)
}