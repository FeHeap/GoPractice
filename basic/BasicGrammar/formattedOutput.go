package main

import "fmt"

func main()  {
	str := "steven"
	fmt.Printf("%T , %v \n", str, str)	//string , steven
	var a rune = '一'
	fmt.Printf("%T , %v \n", a, a)		//int32 , 19968
	var b byte = 'b'
	fmt.Printf("%T , %v \n", b, b)		//uint8 , 98
	var c int32 = 98
	fmt.Printf("%T , %v \n\n", c, c)		//int32 , 98

	var flag bool
	fmt.Printf("%T , %t \n", flag, flag)	//bool , false
	flag = true
	fmt.Printf("%T , %t \n\n", flag, flag)	//bool , true

	fmt.Printf("%T , %d \n", 123, 123)		//int , 123
	fmt.Printf("%T , %5d \n", 123, 123)		//int ,   123
	fmt.Printf("%T , %05d \n", 123, 123)		//int , 00123
	fmt.Printf("%T , %b \n", 123, 123)		//int , 1111011
	fmt.Printf("%T , %o \n", 123, 123)		//int , 173
	fmt.Printf("%T , %c \n", 97, 97)			//int , a
	fmt.Printf("%T , %q \n", 97, 97)			//int , 'a'
	fmt.Printf("%T , %x \n", 123, 123)		//int , 7b
	fmt.Printf("%T , %X \n", 123, 123)		//int , 7B
	fmt.Printf("%T , %U \n\n", '一', '一')	//int32 , U+4E00

	fmt.Printf("%b \n",123.123456)		//8664042977533870p-46
	fmt.Printf("%f \n", 123.1)			//123.100000
	fmt.Printf("%.2f \n", 123.123456)	//123.12
	fmt.Printf("%e \n", 123.123456)		//1.231235e+02
	fmt.Printf("%E \n", 123.123456)		//1.231235E+02
	fmt.Printf("%.1e \n", 123.123456)	//1.2e+02
	fmt.Printf("%F \n", 123.123456)		//123.123456
	fmt.Printf("%g \n", 123.123456)		//123.123456
	fmt.Printf("%G \n\n", 123.123456)	//123.123456

	var value complex64 = 2.1 + 12i
	value2 := complex(2.1, 12)
	fmt.Println(real(value))	//2.1
	fmt.Println(imag(value))	//12
	fmt.Println(value2,"\n")	//(2.1+12i)

	arr := []byte{'x', 'y', 'z', 'z'}
	fmt.Printf("%s \n", "happy blockchain")	//happy blockchain
	fmt.Printf("%q \n", "happy blockchain")	//"happy blockchain"
	fmt.Printf("%x \n", "happy blockchain")	//686170707920626c6f636b636861696e
	fmt.Printf("%X \n", "happy blockchain")	//686170707920626C6F636B636861696E
	fmt.Printf("%T , %s \n", arr, arr)	//[]uint8 , xyzz
	fmt.Printf("%T , %q \n", arr, arr)	//[]uint8 , "xyzz"
	fmt.Printf("%T , %x \n", arr, arr)	//[]uint8 , 78797a7a
	fmt.Printf("%T , %X \n", arr, arr)	//[]uint8 , 78797A7A
}