package main

import (
	"fmt"
	"sync"
)

// channels <=> queues

var wg = sync.WaitGroup{}
var doneCh = make(chan struct{}) // 0 memory footprint, used to detect a message

func main() {
	
	for i:= 0 ; i<5;i++{
		ch := make(chan int, 50) // only way to create a channel, by default channel size is 1
		wg.Add(2)
		go func(ch <- chan int) {	// strange syntax, mean consumer only...
			for j := range ch {	// blocking until detection of the close(ch) or goroutine is stopped
				// can also use for { if j, ok := <- ch; ok {} else{break}}
				fmt.Println(j)
			}
			wg.Done()
		}(ch)
		go func(ch chan <- int, i int) {	// producer only
			ch <- i
			ch <- 42
			close(ch)
			wg.Done()
		}(ch, i)
	}
	wg.Wait()
	// after main all goroutine will be stopped 
	// we could defer func() {close(ch)}() for clean close
}