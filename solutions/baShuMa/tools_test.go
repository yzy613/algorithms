package baShuMa

import "testing"

func TestCantorExpansion(t *testing.T) {
	testArr := []int{3, 4, 1, 5, 2}
	lengthTestArr := len(testArr)
	preprocessFactorialArr(lengthTestArr)
	r := CantorExpansion(testArr)
	if r != 61 {
		t.Errorf("CantorExpansion([3, 4, 1, 5, 2]) failed. Got %v, expected 61.", r)
	}

	testArr = []int{2, 3, 1}
	lengthTestArr = len(testArr)
	preprocessFactorialArr(lengthTestArr)
	r = CantorExpansion(testArr)
	if r != 3 {
		t.Errorf("CantorExpansion([2, 3, 1]) failed. Got %v, expected 3.", r)
	}

	testArr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	lengthTestArr = len(testArr)
	preprocessFactorialArr(lengthTestArr)
	r = CantorExpansion(testArr)
	if r != 0 {
		t.Errorf("CantorExpansion([1, 2, 3, 4, 5, 6, 7, 8, 9]) failed. Got %v, expected 0.", r)
	}

	testArr = []int{6, 5, 2, 3, 1, 4}
	lengthTestArr = len(testArr)
	preprocessFactorialArr(lengthTestArr)
	r = CantorExpansion(testArr)
	if r != 704 {
		t.Errorf("CantorExpansion([6, 5, 2, 3, 1, 4]) failed. Got %v, expected 704.", r)
	}

	testArr = []int{2, 8, 3, 1, 0, 4, 7, 6, 5}
	lengthTestArr = len(testArr)
	preprocessFactorialArr(lengthTestArr)
	r = CantorExpansion(testArr)
	if r != 117485 {
		t.Errorf("CantorExpansion([2, 8, 3, 1, 0, 4, 7, 6, 5]) failed. Got %v, expected 117485.", r)
	}
}

func TestUnCantorExpansion(t *testing.T) {
	testArr := []int{3, 4, 1, 5, 2}
	lengthTestArr := len(testArr)
	preprocessFactorialArr(lengthTestArr)
	temp := UnCantorExpansion(61, lengthTestArr)
	r := CantorExpansion(temp)
	if r != 61 {
		t.Errorf("UnCantorExpansion(61, 5) failed. Got %v, expected [3, 4, 1, 5, 2].", temp)
	}

	testArr = []int{2, 8, 3, 1, 0, 4, 7, 6, 5}
	lengthTestArr = len(testArr)
	preprocessFactorialArr(lengthTestArr)
	temp = UnCantorExpansion(117485, lengthTestArr)
	r = CantorExpansion(temp)
	if r != 117485 {
		t.Errorf("UnCantorExpansion(117485) failed. Got %v, expected [2, 8, 3, 1, 0, 4, 7, 6, 5].", temp)
	}
}
