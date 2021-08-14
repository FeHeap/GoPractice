package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	TestTrim()
	TestTrimFunc()
	TestTrimLeft()
	TestTrimRight()
	TestTrimRightFunc()
	TestTrimSpace()
	TestTrimPrefix()
	TestTrimSuffix()
}

func TestTrim() {
	fmt.Println("TestTrim()")
	fmt.Println(strings.Trim("  Fe chuang  ", " "))
	fmt.Println()
}

func TestTrimFunc() {
	fmt.Println("TestTrimFunc()")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimFunc("_-! & Fe chuang%#$@  ", f))
	fmt.Println()
}

func TestTrimLeft() {
	fmt.Println("TestTrimLeft()")
	fmt.Println(strings.TrimLeft("  Fe chuang  ", " "))
	fmt.Println()
}

func TestTrimRight() {
	fmt.Println("TestTrimRight()")
	fmt.Println(strings.TrimRight("  Fe chuang  ", " "))
	fmt.Println()
}

func TestTrimRightFunc() {
	fmt.Println("TestTrimRightFunc()")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.TrimRightFunc("_-! & Fe chuang%#$@  ", f))
	fmt.Println()
}

func TestTrimSpace() {
	fmt.Println("TestTrimSpace()")
	fmt.Println(strings.TrimSpace(" \t\n suisei is pretty kawaii \n\t\r\n"))
	fmt.Println()
}

func TestTrimPrefix() {
	fmt.Println("TestTrimPrefix()")
	var s = "Hello, World!"
	s = strings.TrimPrefix(s, "Hello")
	fmt.Println(s)
	fmt.Println()
}

func TestTrimSuffix() {
	fmt.Println("TestTrimSuffix()")
	var s = "Hello, goodbye, etc!"
	s = strings.TrimSuffix(s, " goodbye, etc!")
	fmt.Println(s)
	fmt.Println()
}