package main

import (
	"14_Packages/util"
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func main() {
	fmt.Println(util.Length("Hello"))

	xs := []float64{
		12.44, 11.2, 34.5, 1.4, 6.7, 23.4,
	}	
	mean := stat.Mean(xs, nil)
	fmt.Println(mean)

}