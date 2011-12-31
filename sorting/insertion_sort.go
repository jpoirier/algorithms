/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
package sorting


func insert(a []int, pos int, val int) {
	i := pos - 1
	for {
		if i >= 0 && a[i] > val {
			a[i + 1] = a[i]
			i -= 1
			continue
		}
		a[i + 1] = val
		break
	}
}

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		insert(array, i, array[i])
	}
}
