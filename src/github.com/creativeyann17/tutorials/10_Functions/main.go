package main

import (
	"fmt"
)

type Person struct {
	name string
}

type Counter int

func (p Person) toString() {
	fmt.Println("Person/" + p.name)
}

func (c *Counter) inc() {
	*c++
}

func main() {
	for i := 0 ; i<5;i++ {
		sayMessage("Yo!",i)
	}
	sayMessage2("foo", "bar")
	tmp := "bar"
	update(&tmp)
	fmt.Println(tmp)
	fmt.Println(sum(1,2,3,4))

	_, err := divide(1,0)
	if err != nil {
		fmt.Println(err)
	}

	f1 := func(i int) {
		fmt.Println("Anonymous func", i)
	}	//using () at the end of this line will invoke it

	f1(0)
	f1(1)

	callByFunc(1, f1)

	p1 := Person{
		name: "Yann",
	}

	p1.toString()

	var c Counter = 0
	(&c).inc()
	fmt.Println(c)
}

func callByFunc(i int, f func(int)) {
	f(i)
}

func sayMessage2(params ...any) {	// always in the end

	fmt.Println(params...)	// if no ... then params will de displayed as an array
}

func sayMessage(msg string, i int) {
	fmt.Println(msg)
	fmt.Println(i)
}

func update(str *string) {	// pointer = better performance than copy for big objects
	*str = "foo"
}

func sum(integers ... int) (sum int) {	// can pre-declare return value, not often used
	for _, v := range integers {
		sum += v
	}
	return	// if u return a pointer from a stack variable, the compiler will promote the variable to heap
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Divided by zero")
	} 
	return a/b, nil
}