/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
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
	array := []uint{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	CountingSort(array)
	if !is_sorted2(array) {
		t.Fail()
	}
}

func TestHeapSort(t *testing.T) {
	array := []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	HeapSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

func TestInsertionSort(t *testing.T) {
	array := []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	InsertionSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

func TestMedianSort(t *testing.T) {
	array := []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	MedianSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

func TestQuickSort(t *testing.T) {
	array := []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	MedianSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

func TestSelectionSort(t *testing.T) {
	array := []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	SelectionSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}

func TestBucketSort(t *testing.T) {
	array := []int{15, 9, 8, 1, 4, 11, 7, 12, 13, 6, 5, 3, 16, 2, 10, 14}
	BucketSort(array)
	if !is_sorted(array) {
		t.Fail()
	}
}
