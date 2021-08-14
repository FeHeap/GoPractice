package main

import (
	"fmt"
	"strings"
)

func main() {
	TestCompare()
	TestEqualFold()
	TestRepeat()
	TestReplace()
	TestJoin()
}

func TestCompare() {
	fmt.Println("TestCompare()")
	fmt.Println(strings.Compare("abc", "bcd"))
	fmt.Println("abc" < "bcd")
	fmt.Println()
}

func TestEqualFold() {
	fmt.Println("TestEqualFold()")
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println()
}

func TestRepeat() {
	fmt.Println("TestRepeat()")
	fmt.Println("g" + strings.Repeat("o", 8) + "gle")
	fmt.Println()
}

func TestReplace() {
	fmt.Println("TestReplace()")
	fmt.Println(strings.Replace("黑桃二 黑桃三 黑桃四", "黑", "紅", 2))
	fmt.Println(strings.Replace("黑桃二 黑桃三 黑桃四", "黑", "紅", -1))
	fmt.Println()
}

func TestJoin() {
	fmt.Println("TestJoin()")
	s := []string{"abc", "ABC", "123"}
	fmt.Println(strings.Join(s, ","))
	fmt.Println(strings.Join(s, ""))
	fmt.Println()
}