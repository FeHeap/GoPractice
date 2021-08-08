package main

import "fmt"

var a1 int = 7
var b1 int = 9
func main() {
	a1, b1, c1 := 10, 20, 0
	fmt.Printf("In main(), a1 = %d\n", a1) //10
	fmt.Printf("In main(), b1 = %d\n", b1) //20
	fmt.Printf("In main(), c1 = %d\n", c1) //0
	c1 = sum(a1, b1)
	fmt.Printf("In main(), c1 = %d\n", c1) //33
}

func sum(a1, b1 int) (c1 int) {
	a1++
	b1 += 2
	c1 = a1 + b1
	fmt.Printf("In sum(), a1 = %d\n", a1) //11
	fmt.Printf("In sum(), b1 = %d\n", b1) //22
	fmt.Printf("In sum(), c1 = %d\n", c1) //33
	return c1
}