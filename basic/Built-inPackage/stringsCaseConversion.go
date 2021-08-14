package main

import (
	"fmt"
	"strings"
)

func main() {
	TestTitle()
	TestToTitle()
	TestToLower()
	TestToUpper()
}

func TestTitle() {
	fmt.Println("TestTitle()")
	fmt.Println(strings.Title("suisei is very kawaii"))
	fmt.Println()
}

func TestToTitle() {
	fmt.Println("TestToTitle()")
	fmt.Println(strings.ToTitle("suiSei iS VEry KAWAII"))
	fmt.Println()
}

func TestToLower() {
	fmt.Println("TestToLower()")
	fmt.Println(strings.ToLower("suiSei iS VEry KAWAII"))
	fmt.Println()
}

func TestToUpper() {
	fmt.Println("TestToUpper()")
	fmt.Println(strings.ToUpper("suiSei iS VEry KAWAII"))
	fmt.Println()
}

