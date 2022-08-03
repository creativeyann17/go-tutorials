package main

import (
	"fmt"
)

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() (int, error)
}

type WriterCloser interface{
	Writer
	Closer
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {	
	n, err := fmt.Println(string(data))
	return n, err
}

func (cw ConsoleWriter) Close() (int, error) {	
	return 0, nil
}

func main() {
	var w Writer = &ConsoleWriter{}
	var wValue WriterCloser = ConsoleWriter{}
	w.Write([]byte("Yo!"))
	wValue.Write([]byte("Yo2"))
	fmt.Println(w, wValue)
	cw := w.(*ConsoleWriter)	// cast ?
	fmt.Println(cw)
	cw2, err := w.(WriterCloser)	// cast of something it's not
	fmt.Println(cw2, err)
	// when u want to put universal stuff as a param, interface{} or any, anything is nothing
	var i any = 0
	switch i.(type) {
	case int:
		fmt.Println("integer")
	default:
		fmt.Println("something else")
	}
}