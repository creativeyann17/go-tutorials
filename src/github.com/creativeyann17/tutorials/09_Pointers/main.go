package main

import (
	"fmt"
)

type MyStruct struct {
	foo int
}

func main() {
	var a int = 42
	var b *int = &a
	a = 27
	fmt.Println(&a, b)
	fmt.Println(a, *b)
	*b = 42
	fmt.Println(a, *b)

	array := [3]int{1,2,3}
	i:=&array[0]
	j:=&array[1]	// + 8 go don't allow pointer arithmetic or use the 'unsafe' package
	fmt.Printf("%v %p %p\n", array,i,j)

	var s1 *MyStruct
	fmt.Println(s1)
	s1 = &MyStruct{foo: 42}
	s2 := new(MyStruct)
	(*s2).foo=42	//s2.foo seems to work too
	fmt.Println(s1, s2)

	// slices & maps are pointers that's why we had change in previous examples

}