package main

import (
	"fmt"
	"reflect"
)

type Person struct {	// Person will be public
	firstName string 	// private fields
	lastName string
	PublicField string `required max:"100"` // public field
	living // nice to compose structs
}

type living struct {
	age int
}


func main() {
	// Map
	m1 := map[string]int {
		"0": 0,
		"1": 1,
		"2": 2,
	}
	fmt.Println(m1)
	m2 := map[[3]int]int {	// arrays keys =OK, slice = KO, fine
		{0,1,2}: 3,
	}
	fmt.Println(m2)
	m3 := make(map[string]int)
	m3["foo"] = 0
	m3["bar"] = 1	// map key order can change
	m4 := m3	// reference like slice
	delete(m3, "bar")
	delete(m3, "bar")	// ok work
	bar, barOk := m3["bar"]
	_ , fooOk := m3["foo"]
	fmt.Println(m3, fooOk, bar, barOk, len(m3), m4)

	// Structs
	p1 := Person {
		firstName: "Foo",
		lastName: "Bar",
		PublicField: "i'm public and required",
		living: living{age: 42},
	}

	fmt.Println(p1)
	/*p2 := Person{
		"foo", "bar",	// positional instance, ok but pain if add/remove a field later ...
	}*/
	s1 := struct{name string}{name: "foo"}	// anonymous struct
	s2 := s1	// copy
	s2.name = "bar"
	fmt.Println(s1.name)

	t := reflect.TypeOf(Person{})
	field, _ := t.FieldByName("PublicField")
	fmt.Println(field)
}