package main

import (
	"fmt"
)

func main() {
	// arrays, fixed size known at compile time
	a1 := [3]int{1,2,3} 
	a2 := [...]int{1,2}
	var a3 [3]int
	a3[2] = 42
	fmt.Println(a1,a2,a3, a3[2], len(a3), cap(a3))	// arrays will always have len = cap

	t1 := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(t1)
	t2 := t1
	t2[0] = [...]int{7,8,9}	// copy not reference
	fmt.Println(t1, t2)

	// slices, list/seq etc
	s1 := []int{0,1,2,3,4,5,6,7,8,9}
	s2 := s1
	s2[6]=42
	fmt.Println(s1, len(s1), cap(s1), s2)	// copy by reference ><, better to think that slice are mutable object
	fmt.Println(s1[:],s1[2:],s1[:8],s1[2:8])	// work for arrays too

	s3 := make([]int, 3, 100)	// dynamic creation of slice, cap can be bigger than current len
	fmt.Println(s3 , len(s3), cap(s3))

	s4 := []int{}
	s4 = append(s4, 0, 1, 2, 3, 4, 5)
	s4 = append(s4, s4...)	// nice
	fmt.Println(s4, len(s4), cap(s4))

	s5 := []int{0,1,2,3,4,5}	// stack :)
	s6 := s5[1:]
	s7 := s5[:len(s5)-1]
	s8 := append(s5[:2],s5[3:]...) // <!> s5 is scrap like ... shit, really that's a pain, we need immutable slice
	fmt.Println(s5, s6, s7, s8)	


}