package main

import "fmt"

type Phone interface {
	call()
}

type AndroidPhone struct {
}

type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("I'm android phone, and you can take a phone call.")
}

func (i IPhone) call() {
	fmt.Println("I'm IPhone, and you can take a phone call.")
}

func main() {
	var phone Phone
	phone = new(AndroidPhone)
	fmt.Printf("%T , %v , %p \n", phone, phone, &phone)
	phone.call()
	phone = AndroidPhone{}
	fmt.Printf("%T , %v , %p \n", phone, phone, &phone)
	phone.call()
	phone = new(IPhone)
	fmt.Printf("%T , %v , %p \n", phone, phone, &phone)
	phone.call()
	phone = IPhone{}
	fmt.Printf("%T , %v , %p \n", phone, phone, &phone)
	phone.call()
}
