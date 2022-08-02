package main

import (
	"fmt"
)

func main() {
	for i := 0; i< 5 ;i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i/=2
		} else {
			i = 2*i+1
		}
	}
	j :=0
	for j<2 {	// <=> sugar syntax ;j<2;
		j++
	}

	k:=0
	Loop:
	for {
		k++
		if k == 3 {
			break Loop	// like goto
		} else if k ==2 {
			continue
		}
		fmt.Println(k)
	}

	s := []int{1,2,3}
	for k,v := range s {	// range is a keyword, not a function
		fmt.Println(k, v)
	}
}