package math

import (
	"testing"
)

type AddData struct {
	x, y int
	result int
}

func TestAdd(t *testing.T) {

	testData := []AddData{{1, 3, 4},{2, 3, 5}}

	for _, v := range testData {
		res := Add(v.x,v.y)
		if res != v.result {
			t.Errorf("Add(1,2) FAILED")
		} else {
			t.Logf("Add(1,2) SUCCESS")
		}
	}

}