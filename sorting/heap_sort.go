/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
package sorting


func buildHeap(a []int) {
	l := len(a)
	for i := (l/2)-1; i >= 0; i-- {
		heapify(a, i, l)
	}
}

func heapify(a []int, idx int, max int) {
	left := 2 * idx + 1
	right := 2 * idx + 2
	largest := 0
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
		heapify(a, largest, max)
	}
}

func HeapSort(array []int) {
	buildHeap(array)
	for i := len(array)-1; i >= 1; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array, 0, i)
	}
}
