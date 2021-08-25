package main

import "fmt"

type Humann struct {
	name, phone string
	age int
}

type Studentt struct {
	Humann
	school string
}

type Employeee struct {
	Humann
	company string
}

func main() {
	s1 := Studentt{Humann{"Daniel", "15012345***", 13}, "NCHU"}
	e1 := Employeee{Humann{"Steven", "17812345***", 35}, "HW"}

	s1.SayHi()
	e1.SayHi()
}

func (h *Humann) SayHi() {
	fmt.Printf("Hello! I'm %s. I'm %d years old. My phone number is: %s\n", h.name, h.age, h.phone)
}

func (s *Studentt) SayHi() {
	fmt.Printf("Hello! I'm %s. I'm %d years old. I'm studying at %s. My phone number is: %s\n", s.name, s.age, s.school, s.phone)
}

func (e *Employeee) SayHi() {
	fmt.Printf("Hello! I'm %s. I'm %d years old. I work at %s. My phone number is: %s\n", e.name, e.age, e.company, e.phone)
}