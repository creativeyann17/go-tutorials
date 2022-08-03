package main

// goroutines = concurrent dev

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}	// like a join of futures

func main() {
	go sayHello(nil,nil,"test")	// dispatch in a 'green thread' != OS Thread
	wg.Add(2)
	m1 := "foo"
	go func() {	// race condition variable will change after
		sayHello(nil, nil, m1)
		wg.Done()
	}()
	go func(msg string) {	// now we are safe
		sayHello(nil,nil, msg)
		wg.Done()
	}("foo")
	m1 = "bar"
	//time.Sleep(100 * time.Millisecond) // bad but just for testing
	wg.Wait()
	fun()
}

func sayHello(wg *sync.WaitGroup,m* sync.RWMutex, more ... any) {
	fmt.Println(append(more, "Hello!")...)
	if wg !=nil {
		wg.Done()
	}
	if m != nil {
		m.RUnlock()
	}
}

func fun() {
	runtime.GOMAXPROCS(8)	// -1 return current count, by default = nb of cores
	var wg = &sync.WaitGroup{}
	var m = &sync.RWMutex{}
	var counter = 0
	for i := 0 ; i<10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello(wg, m, counter)
		m.Lock()
		go increment(wg, m, &counter)
	}
	wg.Wait()
}

func increment(wg *sync.WaitGroup, m *sync.RWMutex, counter *int) {
	(*counter)++
	m.Unlock()
	wg.Done()
}