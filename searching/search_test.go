/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
package searching

import (
	"testing"
)


func TestBinarySearch(t *testing.T) {
}

func TestBinaryTreeSearch(t *testing.T) {
}

func TestHashSearch(t *testing.T) {
}

func TestSequentialSearch(t *testing.T) {
	array := []int{1, 4, 8, 9, 11, 15, 17, 22, 23, 26}
	s := SequentialSearch(array, 15)
	if !s {
		t.Fail()
	}
	s = SequentialSearch(array, 99)
	if s {
		t.Fail()
	}
}

func TestBinarySearch(t *testing.T) {
	array := []int{1, 4, 8, 9, 11, 15, 17, 22, 23, 26}
	s := BinarySearch(array, 11)
	if !s {
		t.Fail()
	}
	s = BinarySearch(array, 99)
	if s {
		t.Fail()
	}
}
