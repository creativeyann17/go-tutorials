package main

import (
	"fmt"
	"strconv"
)

var (
	firstName = "Yann"
	lastName = "MARCOU"
	i int
	k int
)

var j float32 = 0.123
var PUBLIC = "foo"

func main() {
	println(i, k)
	var i int
	var L int
	var s string
	i = 42
	k := 49
	//l := "foo"
	fmt.Println(i, j ,k, L, s)
	fmt.Printf("%v, %T\n", j, j)
	var j = float64(j)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Println(firstName, lastName, strconv.Itoa(i))

	var b bool = true
	fmt.Printf("%v, %T\n", b, b)
}