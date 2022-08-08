package math

import (
	"testing"
	"time"
)

type AddData struct {
	x, y int
	result int
}

type IsPrimeData struct {
	value int
	result bool
}

func TestAdd(t *testing.T) {

	testData := []AddData{{1, 3, 4},{2, 3, 5}}

	for _, v := range testData {
		res := Add(v.x,v.y)
		if res != v.result {
			t.Errorf("Add(%v) FAILED", v)
		} else {
			t.Logf("Add(%v) SUCCESS", v)
		}
	}

}

func TestIsPrime(t *testing.T) {
	testData := []IsPrimeData{{1, false},{2, false}, {3, true}, {7, true}, {19, true}, {21, false}}

	for _, v := range testData {
		res := isPrime(v.value)
		if res != v.result {
			t.Errorf("IsPrime(%v) FAILED", v)
		} else {
			t.Logf("IsPrime(%v) SUCCESS", v)
		}
	}
}

// saving the result out of the function is more reliable (avoid compiler optimization)
var result int

func BenchmarkAdd(b *testing.B) {
	var r int
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
					r = Add(1,2)
	}
	result = r
}

func BenchmarkSleep(b *testing.B) {
	for n := 0; n < b.N; n++ {
		time.Sleep(100 * time.Millisecond)
	}
}


func BenchmarkIsPrime(b *testing.B) {
	for n := 0; n < b.N / 1000; n++ {
		r := isPrime(n)
		b.Logf("IsPrime(%v) = %v", n, r)
	}
}