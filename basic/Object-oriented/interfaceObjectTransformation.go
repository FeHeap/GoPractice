package main

import (
	"fmt"
	"math"
)

type Shapee interface {
	perimeter() float64
	area() float64
}

type Rectanglee struct {
	a, b float64
}

type Trianglee struct {
	a, b, c float64
}

type Circlee struct {
	radius float64
}

func (r Rectanglee) perimeter() float64 {
	return (r.a + r.b) * 2
}

func (r Rectanglee) area() float64 {
	return r.a * r.b
}

func (t Trianglee) perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Trianglee) area() float64 {
	p := t.perimeter() / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (c Circlee) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circlee) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func getType1(s Shapee) {
	if instance, ok := s.(Rectanglee); ok {
		fmt.Printf("Rectangle: length %.2f , width %.2f , ", instance.a, instance.b)
	} else if instance, ok := s.(Trianglee); ok {
		fmt.Printf("Triangle: edges %.2f , %.2f , %.2f , ", instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Circlee); ok {
		fmt.Printf("Circle: radius %.2f , ", instance.radius)
	}
}

func getType2(s Shapee) {
	switch instance := s.(type) {
	case Rectanglee:
		fmt.Printf("Rectangle: length %.2f , width %.2f , ", instance.a, instance.b)
	case Trianglee:
		fmt.Printf("Triangle: edges %.2f , %.2f , %.2f , ", instance.a, instance.b, instance.c)
	case Circlee:
		fmt.Printf("Circle: radius %.2f , ", instance.radius)
	}
}

func getResult(s Shapee) {
	getType2(s)
	fmt.Printf("Perimeter: %.2f, Area: %.2f\n", s.perimeter(), s.area())
}

func main() {
	var s Shapee
	s = Rectanglee{3, 4}
	getResult(s)
	showInfoo(s)

	s = Trianglee{3, 4, 5}
	getResult(s)
	showInfoo(s)

	s = Circlee{1}
	getResult(s)
	showInfoo(s)

	x := Trianglee{3, 4, 5}
	fmt.Println(x)
}

func (t Trianglee) String() string {
	return fmt.Sprintf("Triangle object, %.2f, %.2f, %.2f", t.a, t.b, t.c)
}

func showInfoo(s Shapee) {
	fmt.Printf("%T, %v\n", s, s)
	fmt.Println("------------------")
}