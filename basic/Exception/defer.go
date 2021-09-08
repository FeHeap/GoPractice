package main

import "fmt"

type person struct {
	firstName, lastName string
}

func (p person) fullName() {
	fmt.Printf("%s %s", p.lastName, p.firstName)
}

func main() {
	a := 5
	b := 6
	defer printAdd(a, b ,true)
	a = 10
	b = 7
	printAdd(a, b, false)
	fmt.Println("-----------")

	str := "suisei きょも かわい"
	fmt.Printf("Origin:\n%s\n", str)
	fmt.Println("Reverse:")
	ReverseString(str)
	fmt.Println("\n-----------")

	defer fmt.Println("-----------")
	p := person{"Fe", "Chuang"}
	defer fmt.Println()
	defer p.fullName()
	defer fmt.Print("Welcome ")

	defer fmt.Println("-----------")
	s := []int{78, 100, 2, 400, 324}
	defer getLargest(s)

	defer fmt.Println("-----------")
	defer funcA()
	funcB()
	defer funcC()
	fmt.Println("main over...")
}

func funcA() {
	fmt.Println("This is funcA")
}

func funcB() {
	fmt.Println("This is funcB")
}

func funcC() {
	fmt.Println("This is funcC")
}

func finished() {
	fmt.Println("Finish!")
}

func getLargest(s []int) {
	defer finished()
	fmt.Println("Start to find the maximum:")
	max := s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	fmt.Printf("The maximum in %v is: %v\n", s, max)
}

func printAdd(a, b int, flag bool) {
	if flag {
		fmt.Printf("Execute printAdd() delay, parameter a = %d, b = %d, a + b = %d\n", a, b, a+b)
	} else {
		fmt.Printf("Execute printAdd(), parameter a = %d, b = %d, a + b = %d\n", a, b, a+b)
	}
}

func ReverseString(str string) {
	for _, v := range []rune(str) {
		defer fmt.Printf("%c", v)
	}
}