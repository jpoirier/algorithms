/*
  
*/
package main

import (
	"fmt"
	"rand"
)

func partition(a []int) int {
	l := len(a)
	left := 0
	right := 0
	pivot := 0
	if l > 0 {
		right = l - 1
	}
	if l < 3 {
		pivot = right
	} else {
		// median of three calculation
		var ps [3]int
		for i := 0; i < 3; i++ {
			ps[i] = int(rand.Uint32() % uint32(l))
		}
		if a[ps[0]] > a[ps[1]] {
			ps[0], ps[1] = ps[1], ps[0]
		}
		if a[ps[1]] > a[ps[2]] {
			ps[1], ps[2] = ps[2], ps[1]
		}
		if a[ps[0]] > a[ps[1]] {
			ps[0], ps[1] = ps[1], ps[0]
		}
		pivot = ps[1]
	}
	a[pivot], a[right] = a[right], a[pivot]
	store := left
	for i := left; i < right; i++ {
		if a[i] <= a[right] {
			a[i], a[store] = a[store], a[i]
			store += 1
		}
	}
	a[store], a[right] = a[right], a[store]
	return store
}

func QuickSort(array []int) {
	if len(array) > 1 {
		// recursively work on the sub-arrays,
		// which are just slices of slices
		pivot := partition(array)
		QuickSort(array[0:pivot])
		QuickSort(array[pivot+1:])
	}
}

func main() {
	array := []int{89, 606, 533, 1999, 3, 1, 22, 604, 605, 77}
	fmt.Println("Unsorted: ", array)
	QuickSort(array)
	fmt.Println("Sorted  : ", array)
	for i := len(array) - 1; i > 0; i-- {
		if array[i] < array[i-1] {
			fmt.Println("FAILED")
			return
		}
	}
	fmt.Println("PASSED")
}
