package main

import (
	"fmt"
	"regexp"
)

func main() {
	TestRegexp()
}

func TestRegexp() {
	//Match(pattern string, b []byte) (matched bool, err error)
	flag, _ := regexp.Match("^\\d{6,7}$", []byte("123456789"))
	fmt.Println(flag)
	fmt.Println()

	//MatchString(pattern string, s string) (matched bool, err error)
	flag, _ = regexp.MatchString("^\\d{6,7}$", "0123456789")
	fmt.Println(flag)
	fmt.Println()

	//Compile(expr string) (*Regexp, error)
	Regexp1, _ := regexp.Compile("^\\d{6}\\D{2}$")

	//MustCompile(str string) *Regexp
	Regexp2 := regexp.MustCompile("^[\u4e00-\u9fa5]+$")

	//Match(b []byte) bool
	flag = Regexp2.Match([]byte("星街彗星"))
	fmt.Println("xxx:", flag)
	fmt.Println()

	//MatchString(s string) bool
	flag = Regexp1.MatchString("123456ab")
	fmt.Println(flag)
	fmt.Println()

	//ReplaceAll(src, repl []byte) []byte
	text := "將字符串 123 按照正規表達式 3 4 5 分割成子字串 56 78 組成切片"
	Regexp3 := regexp.MustCompile("[\\d\\s]+")
	text = string(Regexp3.ReplaceAll([]byte(text), []byte("-")))
	fmt.Println("After replacing, text: ", text)
	fmt.Println()

	//Split(s string, n int) []string
	text = "第一部分#第二部分##第三部分###第四部分#第五部分##第六部分"
	MyRegexp := regexp.MustCompile("#+")
	arr := MyRegexp.Split(text, 5)
	fmt.Println(arr)
}
