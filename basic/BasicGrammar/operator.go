package main

import "fmt"

func main()  {
	var a int = 21
	var b int = 10
	var c int
	fmt.Printf("a = %d, b = %d\n", a, b)
	c = a + b
	fmt.Printf("c = a + b, c = %d\n", c)
	c = a - b
	fmt.Printf("c = a - b, c = %d\n", c)
	c = a * b
	fmt.Printf("c = a * b, c = %d\n", c)
	c = a / b
	fmt.Printf("c = a / b, c = %d\n", c)
	c = a % b
	fmt.Printf("c = a %% b, c = %d\n", c)
	a++
	fmt.Printf("a++, a = %d\n", a)
	a = 21
	a--
	fmt.Printf("a = 21, a--, a = %d\n\n", a)

	a = 21
	b = 10
	fmt.Printf("a = %d, b = %d\n", a, b)
	if(a == b) {
		fmt.Println("a == b")
	} else {
		fmt.Println("a != b")
	}
	if(a < b){
		fmt.Println("a < b")
	} else {
		fmt.Println("!(a < b)")
	}
	if(a > b){
		fmt.Println("a > b")
	} else {
		fmt.Println("!(a > b)")
	}
	a = 5
	b = 20
	fmt.Printf("a = %d, b = %d\n", a, b)
	if(a <= b){
		fmt.Println("a <= b")
	}
	if(b >= a){
		fmt.Println("b >= a\n")
	}

	var x bool = true
	var y bool = false
	fmt.Printf("x = %t, y = %t\n", x, y)
	if(x && y){
		fmt.Println("x && y == true")
	}
	if(x || y){
		fmt.Println("x || y == true")
	}
	x = false
	y = true
	fmt.Printf("x = %t, y = %t\n", x, y)
	if(x && y){
		fmt.Println("x && y == true")
	} else {
		fmt.Println("x && y == false")
	}
	if(!(x && y)){
		fmt.Println("!(x && y) == true\n")
	}

	var m uint = 60		/* 60 = 0011 1100 */
	var n uint = 13		/* 13 = 0000 1101 */
	var o uint = 0
	fmt.Printf("m = %d, n = %d\n", m, n)
	o = m & n			/* 12 = 0000 1100 */
	fmt.Printf("o = m & n, o = %d\n", o)
	o = m | n			/* 61 = 0011 1101 */
	fmt.Printf("o = m | n, o = %d\n", o)
	o = m ^ n			/* 49 = 0011 0001 */
	fmt.Printf("o = m ^ n, o = %d\n", o)
	o = m << 2		   /* 240 = 1111 0000 */
	fmt.Printf("o = a << 2, o = %d\n", o)
	o = m >> 2			/* 15 = 0000 1111 */
	fmt.Printf("o = a >> 2, o = %d\n\n",o)

	var s int = 21
	var t int
	fmt.Printf("s = %d\n", s)
	t = s
	fmt.Printf("t = s, t = %d\n", t)
	t += s
	fmt.Printf("t += s, t = %d\n", t)
	t -= s
	fmt.Printf("t -= s, t = %d\n", t)
	t *= s
	fmt.Printf("t *= s, t = %d\n", t)
	t /= s
	fmt.Printf("t /= s, t = %d\n", t)
	t = 200
	fmt.Printf("t = %d\n", t)
	t <<= 2
	fmt.Printf("t <<= 2, t = %d\n", t)
	t >>= 2
	fmt.Printf("t >>= 2, t = %d\n", t)
	t &= 2
	fmt.Printf("t &= 2, t = %d\n", t)
	t ^= 2
	fmt.Printf("t ^= 2, t = %d\n", t)
	t |= 2
	fmt.Printf("t |= 2, t = %d\n\n", t)

	var i int = 4
	var j int32
	var k float32
	var ptr *int
	fmt.Printf("type of i = %T\n", i)
	fmt.Printf("type of j = %T\n", j)
	fmt.Printf("type of k = %T\n", k)
	ptr = &i
	fmt.Printf("i = %d\n", i)
	fmt.Printf("*ptr = %d\n", *ptr)
}