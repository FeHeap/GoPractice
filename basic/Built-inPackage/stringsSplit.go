package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	TestFields()
	TestFieldsFunc()
	TestSplit()
	TestSplitN()
	TestSplitAfter()
	TestSplitAfterN()
}

func TestFields() {
	fmt.Println("TestFields()")
	fmt.Println(strings.Fields(" abc 123 ABC xyz XYZ"))
	fmt.Println()
}

func TestFieldsFunc() {
	fmt.Println("TestFieldsFunc()")
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Println(strings.FieldsFunc("abc@123*&xyz%XYZ", f))
	fmt.Println()
}

func TestSplit() {
	fmt.Println("TestSplit()")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	fmt.Println()
}

func TestSplitN() {
	fmt.Println("TestSplitN()")
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 1))
	fmt.Println()
}

func TestSplitAfter() {
	fmt.Println("TestSplitAfter()")
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	fmt.Println()
}

func TestSplitAfterN() {
	fmt.Println("TestSplitAfterN()")
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",",2))
	fmt.Println()
}