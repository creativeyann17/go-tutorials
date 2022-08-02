package main

import (
	"fmt"
)

const myConst int8 = 27 // will be shadowed
const (
	a = iota
	b	// will repeat iota
	c // will repeat iota
)
const (
	_ = iota	// like scala, we dont care about it
	d	// reset of the value to 0, nice to create enums
)
const (
	_ = 1 << iota	// just a dream for return code/errors <3
	enum1
	enum2
	enum3
)

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func main() {
	const myConst = 42	// MyConst if public
	var v int16 = 27 // compiler infer type
	fmt.Printf("%v %T\n", myConst + v, myConst + v)	// can't be done with variable, myConst will be of type int16 by the compiler

	fmt.Printf("%v %T\n", a, a)	
	fmt.Printf("%v %T\n", b, b)	
	fmt.Printf("%v %T\n", c, c)	
	fmt.Printf("%v %T\n", d, d)	

	var test = d
	fmt.Println(test == d)

	fmt.Println(enum1, enum2, enum3)
	fmt.Println(KB, MB, GB)
	fmt.Printf("%.2fGB\n", 400000000.0/GB)

	var roles = enum1 | enum3
	fmt.Printf("%b %v\n", roles, roles & enum1 > 0)
}
