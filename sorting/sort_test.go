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

func is_sorted2(a []uint) bool {
	for i := len(a) - 1; i > 0; i-- {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

func TestCountingSort(t *testing.T) {
	array := []uint{19, 2, 23, 0, 24, 7, 11, 22, 4, 21, 21, 2}
	CountingSort(array)
	if !is_sorted2(array) {
		t.Fail()
	}
}

func TestHeapSort(t *testing.T) {
	array := []int{88, 54, 37, 28, 9, 2, 1, 7, 19, 2, 23, 0, 24, 7, 11, 22, 4, 21, 21, 52}
	HeapSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
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

func TestQuickSort(t *testing.T) {
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


//func TestBucketSort(t *testing.T) {
//	array := []int{89, 606, 533, 1999, 3, 1, 22, 604, 605, 77}
//	QuickSort(array)
//	if !is_sorted(array) {
//		t.Fail()
//	}
//}
