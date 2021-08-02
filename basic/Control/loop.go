package main

import "fmt"

func main()  {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ",i)
	}
	fmt.Println()

	i := 0
	for ; i <= 10; i++ {
		fmt.Printf("%d ",i)
	}
	fmt.Println()
	i = 0
	for ; ; i++ {
		if(i > 20) {
			break
		}
		fmt.Printf("%d ",i)
	}
	fmt.Println()

	var j int
	for j <= 10 {
		fmt.Print(j)
		j++
	}
	fmt.Println()

	var k int
	for {
		if k > 10 {
			break
		}
		fmt.Print(k)
		k++
	}
	fmt.Println()

	str := "123ABCabc一丁ㄎ"
	for i, value := range str {
		fmt.Printf("%dth's ASCII: %d, char is %c \n", i, value, value)
	}

	sum := 0
	for i:= 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	i = 1
	sum = 0
	for i <= 40 {
		if i % 3 == 0 {
			sum += i
			fmt.Print(i)
			if i < 39 {
				fmt.Print("+")
			} else {
				fmt.Printf(" = %d \n", sum)
			}
		}
		i++
	}

	count := 0
	for n := 32.0; n >= 4; n -= 1.5 {
		count++
	}
	fmt.Println(count)

	lines := 8
	for n := 0; n < lines; n++ {
		for p := 0; p < 2*n+1; p++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for n := 1; n <= 9; n++ {
		for p := 1; p <= n; p++ {
			fmt.Printf("%d*%d=%d ", p, n, n*p)
		}
		fmt.Println()
	}

	fmt.Print("Prime number 1-50: ")
	var a, b int
	for a = 2; a <= 50; a++ {
		for b = 2; b <= 50; b++ {
			if a % b == 0 {
				break
			}
		}
		if b > (a / b) {
			fmt.Printf("%d\t",a)
		}
	}
	fmt.Println()

	for n := 1; n <= 10; n++ {
		if n > 5 {
			break
		}
		fmt.Printf("%d ",n)
	}
	fmt.Println("\nline after for loop")

	for n := 1; n <= 10; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	var C, c int
	C = 1
	LOOP:
		for C < 50 {
			C++
			for c = 2; c < C; c++ {
				if C % c == 0 {
					goto LOOP
				}
			}
			fmt.Printf("%d \t", C)
		}
}
