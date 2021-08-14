package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	TestContains()
	TestContainsAny()
	TestContainsRune()
	TestCount()
	TestHasPrefix()
	TestHasSuffix()
	TestIndex()
	TestIndexAny()
	TestIndexByte()
	TestIndexRune()
	TestIndexFunc()
	TestLastIndex()
	TestLastIndexAny()
	TestLastIndexByte()
	TestLastIndexFunc()
	res := GetFileSuffix("abc.xyz.lmn.jpg")
	fmt.Println(res)
}

func TestContains() {
	fmt.Println("TestContains()")
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("星街彗星 is my wife.", "彗"))
	fmt.Println()
}

func TestContainsAny() {
	fmt.Println("TestContainsAny()")
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
	fmt.Println()
}

func TestContainsRune() {
	fmt.Println("TestContainsRune()")
	fmt.Println(strings.ContainsRune("星街彗星", '街'))
	fmt.Println(strings.ContainsRune("星街彗星", 34903))
	fmt.Println()
}

func TestCount() {
	fmt.Println("TestCount()")
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("one", ""))
	fmt.Println()
}

func TestHasPrefix() {
	fmt.Println("TestHasPrefix()")
	fmt.Println(strings.HasPrefix("1000phone news", "1000"))
	fmt.Println(strings.HasPrefix("1000phone news", "1000a"))
	fmt.Println()
}

func TestHasSuffix() {
	fmt.Println("TestHasSuffix()")
	fmt.Println(strings.HasSuffix("1000phone news", "news"))
	fmt.Println(strings.HasSuffix("1000phone news", "new"))
	fmt.Println()
}

func TestIndex() {
	fmt.Println("TestIndex()")
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	fmt.Println()
}

func TestIndexAny() {
	fmt.Println("TestIndexAny()")
	fmt.Println(strings.IndexAny("xyzABC123", "鯊鯊: A"))
	fmt.Println()
}

func TestIndexByte() {
	fmt.Println("TestIndexByte()")
	fmt.Println(strings.IndexByte("xyzABC123", 'A'))
	fmt.Println()
}

func TestIndexRune() {
	fmt.Println("TestIndexRune()")
	fmt.Println(strings.IndexRune("abcABC120", 'C'))
	fmt.Println(strings.IndexRune("abcABC120", 'Z'))
	fmt.Println(strings.IndexRune("IT培訓教育", '教'))
	fmt.Println()
}

func TestIndexFunc() {
	fmt.Println("TestIndexFunc()")
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 台灣!", f))
	fmt.Println()
}

func TestLastIndex() {
	fmt.Println("TestLastIndex()")
	fmt.Println(strings.LastIndex("Steven learn english", "e"))
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
	fmt.Println()
}

func TestLastIndexAny() {
	fmt.Println("TestLastIndexAny()")
	fmt.Println(strings.LastIndexAny("chicken", "aiueoy"))
	fmt.Println(strings.LastIndexAny("crwth", "aiueoy"))
	fmt.Println()
}

func TestLastIndexByte() {
	fmt.Println("TestLastIndexByte()")
	fmt.Println(strings.LastIndexByte("abcABCA123", 'A'))
	fmt.Println()
}

func TestLastIndexFunc() {
	fmt.Println("TestLastIndexFunc()")
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.LastIndexFunc("Hello, 世界!", f))
	fmt.Println(strings.LastIndexFunc("Hello, world星詠者", f))
	fmt.Println()
}

func GetFileSuffix(str string) string {
	arr := strings.Split(str, ".")
	return arr[len(arr) - 1]
}