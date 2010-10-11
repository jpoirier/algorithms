// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort

import (
//	"fmt"
	"testing"
)


func is_sorted(a []int) bool {
	for i := len(a) - 1; i > 0; i-- {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

func TestCountingSort(t *testing.T) {
//	k := 25 // the array can contain values ranging from 0 to 24
	array := []int{19, 2, 23, 0, 24, 7, 11, 22, 4, 21, 21, 2}
	CountingSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
//	if string(array) != string(orig) {
//		fmt.Println(string(orig))
////		t.Fail("Unsorted: %s \n Sorted: %s", orig, array)
//		t.Failed()
//	}
}

func TestInsertionSort(t *testing.T) {
	array := []int{88, 54, 37, 28, 9, 2, 1}
	InsertionSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

func TestMedianSort(t *testing.T) {
	array := []int{89, 606, 533, 1999, 3, 1, 22, 604, 605, 77}
	MedianSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

func TestSelectionSort(t *testing.T) {
	array := []int{19, 2, 23, 0, 24, 7, 11, 22, 4, 21, 21, 2}
	SelectionSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

func TestHeapSort(t *testing.T) {
	array := []int{88, 54, 37, 28, 9, 2, 1, 7}
	HeapSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

//func TestBucketSort(t *testing.T) {
//	array := []int{89, 606, 533, 1999, 3, 1, 22, 604, 605, 77}
//	QuickSort(array)
//	if !is_sorted(array) {
//		t.Fail()
//	}
//}



