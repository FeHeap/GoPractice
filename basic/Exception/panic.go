package main

import "fmt"

func TestA() {
	fmt.Println("func TestA()")
}

func TestB() {
	panic("func TestB(): panic")
}

func TestC() {
	fmt.Println("func TestC()")
}

func TestD(x int) {
	var a [100]int
	a[x] = 1000
}

func main() {
	TestA()
	TestD(101)
	TestB()
	TestC()
}