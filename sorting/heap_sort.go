/*

*/
package main

import (
	"fmt"
)

func buildHeap(a []int) {
	l := len(a)
	for i := (l/2)-1; i >= 0; i-- {
		heapify(a[i:l])
	}
}

func heapify(a []int) {
	idx := 0
	left := 1
	right := 2
	largest := 0
	max := len(a)-1
	if left < max && a[left] > a[idx] {
		largest = left
	} else {
		largest = idx
	}
	if right < max && a[right] > a[largest] {
		largest = right
	}
	if largest != idx {
		a[idx], a[largest] = a[largest], a[idx]
		heapify(a[largest:max+1])
	}
}

func HeapSort(array []int) {
	buildHeap(array)
	for i := len(array)-1; i >= 1; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array[0:i+1])
	}
}

func main() {
	array := []int{88, 54, 37, 28, 9, 2, 1, 7}
	fmt.Println("Unsorted: ", array)
	HeapSort(array)
	fmt.Println("Sorted  : ", array)
	for i := len(array) - 1; i > 0; i-- {
		if array[i] < array[i-1] {
			fmt.Println("FAILED")
			return
		}
	}
	fmt.Println("PASSED")
}
