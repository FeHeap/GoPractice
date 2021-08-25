package main

import "fmt"

type Human_ struct {
	name, phone string
	age int
}

type Student_ struct {
	Human_
	school string
}

type Employee_ struct {
	Human_
	company string
}

func main() {
	s1 := Student_{Human_{"Daniel", "15012345***", 13}, "NCHU"}
	e1 := Employee_{Human_{"Steven", "17812345***", 35}, "HW"}
	s1.SayHi()
	e1.SayHi()
}

func (h *Human_) SayHi() {
	fmt.Printf("Hello! I'm %s. I'm %d years old. My phone number is: %s\n", h.name, h.age, h.phone)
}