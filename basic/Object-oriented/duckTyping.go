package main

import "fmt"

type ISayHello interface {
	SayHello() string
}

type DDuck struct {
	name string
}

type PPerson struct {
	name string
}

func (d DDuck) SayHello() string {
	return d.name + ": ga ga ga!"
}

func (p PPerson) SayHello() string {
	return p.name + ": Hello!"
}

func main() {
	duck := DDuck{"Yaya"}
	person := PPerson{"Fe"}
	fmt.Println(duck.SayHello())
	fmt.Println(person.SayHello())
	fmt.Println("------------------")

	var i ISayHello
	i = duck
	fmt.Printf("%T , %v , %p \n", i, i, &i)
	fmt.Println(i.SayHello())
	i = person
	fmt.Printf("%T , %v , %p \n", i, i, &i)
	fmt.Println(i.SayHello())
}