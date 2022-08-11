package main

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
	"github.com/logrusorgru/aurora"
)

func main() {

writer := uilive.New()
writer2 := writer.Newline()
writer3 := writer.Newline()
// start listening for updates and render
writer.Start()

for i := 0; i <= 100; i++ {
  //fmt.Fprintf(writer, "Downloading.. 1 (%d/%d) GB\n", i, 100)
	fmt.Fprintf(writer2, "Downloading.. 2 (%d/%d) GB\n", i, 100)
	fmt.Fprintf(writer3, "Downloading.. 3 (%d/%d) GB\n", i, 100)
  time.Sleep(time.Millisecond * 20)
}

fmt.Fprintln(writer, aurora.Bold("Finished: Downloaded 100GB"))
writer.Stop() // flush and stop rendering
}