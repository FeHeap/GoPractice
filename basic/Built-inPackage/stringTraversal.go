package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我愛Go語言"
	fmt.Println("字節長度:", len(s))
	fmt.Println("--------------")


	//for ... range traverse string
	len := 0
	for i, ch := range s {
		fmt.Printf("%d:%c ", i, ch)
		len++
	}
	fmt.Println()
	fmt.Println("字符串長度", len)
	fmt.Println("--------------")


	//traverse all the bytes
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x ", i, ch)
	}
	fmt.Println()
	fmt.Println("--------------")


	//traverse all the chars
	count := 0
	for i, ch := range []rune(s) {
		fmt.Printf("%d:%c ", i, ch)
		count++
	}
	fmt.Println()
	fmt.Println("字符串長度", count)
	fmt.Println("字符串長度", utf8.RuneCountInString(s))
}
