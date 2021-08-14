package main

import (
	"fmt"
	"strconv"
)

func main() {
	TestItoa()
	TestFormatInt()
	TestFormatUint()
	TestFormatFloat()
}

func TestItoa() {
	fmt.Println("TestItoa()")
	s := strconv.Itoa(199)
	fmt.Printf("%T, %v, len: %d\n", s, s, len(s))
	fmt.Println()
}

func TestFormatInt() {
	fmt.Println("TestFormatInt()")
	s := strconv.FormatInt(-19968, 16)
	fmt.Printf("%T, %v, len: %d\n", s, s, len(s))
	s = strconv.FormatInt(-40869, 16)
	fmt.Printf("%T, %v, len: %d\n", s, s, len(s))
	fmt.Println()
}

func TestFormatUint() {
	fmt.Println("TestFormatUint()")
	s := strconv.FormatInt(19968, 16)
	fmt.Printf("%T, %v, len: %d\n", s, s, len(s))
	s = strconv.FormatInt(40869, 16)
	fmt.Printf("%T, %v, len: %d\n", s, s, len(s))
	fmt.Println()
}

func TestFormatFloat() {
	fmt.Println("TestFormatFloat()")
	s := strconv.FormatFloat(3.1415926, 'g', -1, 64)
	fmt.Printf("%T, %v, len: %d\n", s, s, len(s))
	fmt.Println()
}