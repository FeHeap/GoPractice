package main

import "fmt"

const (
	Unknown = 0
	Female = 1
	Male = 2
)

const (
	a = 10
	b
	c
)

const (
	A = iota
	B = iota
	C = iota
)

const (
	x = iota
	y
	z
)

func main()  {
	fmt.Println(Unknown, Female, Male)	//0 1 2
	fmt.Println(a, b, c)	//10 10 10
	fmt.Println(A, B, C)	//0 1 2
	fmt.Println(x, y, z)	//0 1 2
}