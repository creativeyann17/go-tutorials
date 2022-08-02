package main

import (
	"fmt"
)

func main() {
	var b bool
	n := 1 == 1
	m := 2 == 2
	fmt.Println(b, n, m)
	fmt.Printf("%v, %T\n", b, b)
	number := 42
	fmt.Printf("%v, %T\n", number, number)
	var unsigned uint32 = 42
	fmt.Printf("%v, %T\n", unsigned, unsigned)
	var signed int64 = 42
	fmt.Printf("%v, %T\n", signed, signed)
	a := 10	// 1010
	c := 3  // 0011
	fmt.Println(a + c)
	fmt.Println(a - c)
	fmt.Println(a * c)
	fmt.Println(a / c)
	fmt.Println(a % c)

	fmt.Println(a & c)
	fmt.Println(a | c)
	fmt.Println(a ^ c)
	fmt.Println(a &^ c)

	fmt.Println(a >> 3)
	fmt.Println(a << 3)

	var intNormal int = 10
	var int8 int8 = 3
	fmt.Println(intNormal + int(int8))

	var float1 float32  = 3.14
	var float2 float64 = 13.7e72	// to big for float32
	var float3 float32 = 2.1E14

	fmt.Println(float1, float2, float3)
	decimal := (float3/float1)
	fmt.Printf("%v, %T\n", decimal, decimal)

	var i complex64 = 1 + 2i	// <=> complex(1, 2)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", real(i), real(i))
	fmt.Printf("%v, %T\n", imag(i), imag(i))

	s := "this is " + "a string"	// immutable yeahhhh
	array := []byte(s)
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", s[1], s[1])
	fmt.Printf("%v %T\n", array, array)

	r := 'a'	// it's a rune (not a char) UTF32 encoded
	fmt.Printf("%v %T\n", r, r)

}