package main

import "fmt"

type Teacher struct {
	name string
	age int
	married bool
	sex int8
}

func main() {
	//int
	a := 10
	fmt.Printf("1. The address of a: %p, Value: %v\n\n", &a, a)
	fmt.Printf("====The address of type-int variable a: %p\n\n", a)
	changeIntVal(a)
	fmt.Printf("2. After changeIntVal(a), the address of a: %p, Value: %v\n\n", &a, a)
	changeIntPtr(&a)
	fmt.Printf("3. After changeIntPtr(&a), the address of a: %p, Value: %v\n\n", &a, a)
	fmt.Println("\n")


	//slice
	b := []int{1, 2, 3, 4}
	fmt.Printf("1. The address of b: %p, Value: %v\n\n", &b, b)
	fmt.Printf("====The address of type-slice variable b: %p\n\n", b)
	changeSliceVal(b)
	fmt.Printf("2. After changeSliceVal(b), the address of b: %p, Value: %v\n\n", &b, b)
	changeSlicePtr(&b)
	fmt.Printf("3. After changeSlicePtr(&b), the address of b: %p, Value: %v\n\n", &b, b)
	fmt.Println("\n")


	//array
	c := [4]int{1, 2, 3, 4}
	fmt.Printf("1. The address of c: %p, Value: %v\n\n", &c, c)
	fmt.Printf("====The address of type-array variable c: %p\n\n", c)
	changeArrayVal(c)
	fmt.Printf("2. After changeArrayVal(c), the address of c: %p, Value: %v\n\n", &c, c)
	changeArrayPtr(&c)
	fmt.Printf("3. After changeArrayPtr(&c), the address of c: %p, Value: %v\n\n", &c, c)
	fmt.Println("\n")


	//struct
	d := Teacher{"Fe", 20, false, 1}
	fmt.Printf("1. The address of d: %p, Value: %v\n\n", &d, d)
	fmt.Printf("====The address of type-struct variable d: %p\n\n", d)
	changeStructVal(d)
	fmt.Printf("2. After changeStructVal(d), the address of d: %p, Value: %v\n\n", &d, d)
	changeStructPtr(&d)
	fmt.Printf("3. After changeStructPtr(&d), the address of d: %p, Value: %v\n\n", &d, d)
}

func changeIntVal(a int) {
	fmt.Printf("----In func changeIntVal(a int), the address of a: %p, value = %v\n", &a, a)
	a = 90
}

func changeIntPtr(a *int) {
	fmt.Printf("----In func changeIntPtr(a int), the address of a: %p, value = %v\n", &a, a)
	*a = 40
}

func changeSliceVal(b []int) {
	fmt.Printf("----In func changeIntVal(b []int), the address of b: %p, value = %v\n", &b, b)
	fmt.Printf("----In func changeSliceVal(b []int), the address of type-slice variable b: %p\n", b)
	b[0] = 99
}

func changeSlicePtr(b *[]int) {
	fmt.Printf("----In func changeIntPtr(b []int), the address of b: %p, value = %v\n", &b, b)
	(*b)[1] = 250
}

func changeArrayVal(c [4]int) {
	fmt.Printf("----In func changeArrayVal(c [4]int), the address of c: %p, value = %v\n", &c, c)
	fmt.Printf("----In func changeArrayVal(c [4]int), the address of type-array variable c: %p\n", c)
	c[0] = 99
}

func changeArrayPtr(c *[4]int) {
	fmt.Printf("----In func changeArrayPtr(c *[4]int), the address of c: %p, value = %v\n", &c, c)
	(*c)[1] = 250
}

func changeStructVal(d Teacher) {
	fmt.Printf("----In func changeStructVal(d Teacher), the address of d: %p, value = %v\n", &d, d)
	fmt.Printf("----In func changeStructVal(d Teacher), the address of type-struct variable d: %p\n", d)

	d.name = "Zoe"
	d.sex = 0
}

func changeStructPtr(d *Teacher) {
	fmt.Printf("----In func changeStructPtr(d *Teacher), the address of d: %p, value = %v\n", &d, d)
	d.name = "Luder"
	d.sex = 0
	d.age = 21
}