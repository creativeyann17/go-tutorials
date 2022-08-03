package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("1")
	defer fmt.Println("2")	// call before end of the current block like :after in css
	fmt.Println("3")

	a := "start"
	defer fmt.Println(a)
	a = "end"

	robots()
	//server()	// try to start the program twice and it will trigger the panic

	defer func() {	// will 'save' the state of params u pass to the func
		if err := recover(); err !=nil {
			log.Println("Error: ", err)
			panic(err)	// if u don't know what to do with that error <=> re-throw
		}
	}()

	i, j := 1,0
	ans := i/j
	fmt.Println(ans)

	// panic("my custom panic message")
	// panic are called after defer
}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r* http.Request){
		w.Write([]byte("Hello World!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err !=nil {
		panic(err.Error())
	}
}

func robots() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err!= nil {
		log.Fatal(err)
	}
	fmt.Println(string(robots))
}