package main

import "fmt"

type Student struct {
	name string
	age int
	married bool
	sex int8
}

const COUNT int = 4

func main() {
	a := 10
	fmt.Printf("address of the variable: %x\n\n", &a)


	var b int = 120
	var ip *int
	ip = &b
	fmt.Printf("The type of b: %T, Value: %v\n", b, b)
	fmt.Printf("The type of &b: %T, Value: %v\n", &b, &b)
	fmt.Printf("The type of ip: %T, Value: %v\n", ip, ip)
	fmt.Printf("The type of *ip: %T, Value: %v\n", *ip, *ip)
	fmt.Printf("The type of *&b: %T, Value: %v\n", *&b, *&b)
	fmt.Println(b, &b, *&b)
	fmt.Println(ip, &ip, *ip, *(&ip), &(*ip))
	fmt.Println()


	var s1 = Student{"Steven", 35, true, 1}
	var s2 = Student{"Sunny", 20, false, 0}
	var c *Student = &s1
	var d *Student = &s2
	fmt.Printf("The type of s1: %T, Value: %v\n", s1, s1)
	fmt.Printf("The type of s2: %T, Value: %v\n", s2, s2)
	fmt.Printf("The type of c: %T, Value: %v\n", c, c)
	fmt.Printf("The type of d: %T, Value: %v\n", d, d)
	fmt.Printf("The type of *c: %T, Value: %v\n", *c, *c)
	fmt.Printf("The type of *d: %T, Value: %v\n", *d, *d)
	fmt.Println(s1.name, s1.age, s1.married, s1.sex)
	fmt.Println(c.name, c.age, c.married, c.sex)
	fmt.Println(s2.name, s2.age, s2.married, s2.sex)
	fmt.Println(d.name, d.age, d.married, d.sex)
	fmt.Println((*c).name, (*c).age, (*c).married, (*c).sex)
	fmt.Println((*d).name, (*d).age, (*d).married, (*d).sex)
	fmt.Printf("The type of &c: %T, Value: %v\n", &c, &c)
	fmt.Printf("The type of &d: %T, Value: %v\n", &d, &d)
	fmt.Println(&c.name, &c.age, &c.married, &c.sex)
	fmt.Println(&d.name, &d.age, &d.married, &d.sex)
	fmt.Println()


	e := 3158
	f := &e
	fmt.Println("The address of e:", f)
	fmt.Println("The value of *f:", *f)
	*f++
	fmt.Println("The new value of e:", e)
	fmt.Println()


	g := 58
	fmt.Println("Before func change(), the value of g:", g)
	fmt.Printf("%T\n", g)
	fmt.Printf("%x\n", &g)
	var h *int = &g
	change(h)
	fmt.Println("After func change(), the value of g:", g)
	fmt.Println()

	k := [COUNT]string{"abc", "ABC", "123", "one two three"}
	var ptr [COUNT]*string
	fmt.Printf("%T, %v\n", ptr, ptr)
	for i := 0; i < COUNT; i++ {
		ptr[i] = &k[i]
	}
	fmt.Printf("%T, %v\n", ptr, ptr)
	fmt.Println(ptr[0])
	for i := 0; i < COUNT; i++ {
		fmt.Printf("k[%d] = %s\n", i, *ptr[i])
	}
	fmt.Println()

	var m int
	var ptrm *int
	var pptrm **int
	m = 1234
	ptrm = &m
	fmt.Println("ptrm", ptrm)
	pptrm = &ptrm
	fmt.Println("pptrm", pptrm)
	fmt.Printf("Variable m = %d\n", m)
	fmt.Printf("Pointer variable *ptrn = %d\n", *ptrm)
	fmt.Printf("Pointer variable pointing to pointer **pptrm = %d\n", **pptrm)
}

func change(val *int) {
	*val = 15
}