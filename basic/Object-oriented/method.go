package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func main() {
	r1 := Rectangle{10, 4}
	r2 := Rectangle{12, 5}
	c1 := Circle{1}
	c2 := Circle{10}
	fmt.Println("The area of r1:", r1.Area())
	fmt.Println("The area of r2:", r2.Area())
	fmt.Println("The area of c1:", c1.Area())
	fmt.Println("The area of c2:", c2.Area())
	fmt.Println("--------------------")

	r3 := Rectangle{5, 8}
	r4 := r3
	fmt.Printf("The address of r3: %p\n", &r3)
	fmt.Printf("The address of r4: %p\n", &r4)
	r3.setValue1()
	fmt.Println("r3.height =", r3.height)	//8
	fmt.Println("r4.height =", r4.height)	//8
	fmt.Println("--------------------")
	r3.setValue2()
	fmt.Println("r3.height =", r3.height)	//8
	fmt.Println("r4.height =", r4.height)	//20
	fmt.Println("--------------------")
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) setValue1() {
	fmt.Printf("In setValue1(), The address of r: %p\n", &r)
	r.height = 10
}

func (r *Rectangle) setValue2() {
	fmt.Printf("In setValue2(), The address of r: %p\n", r)
	r.height = 20
}