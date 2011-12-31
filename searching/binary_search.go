/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
package searching


func BinarySearch(array []int, t int) bool {
	low := 0
	high := len(array) - 1
	for {
		if low <= high {
			ix := (low + high) / 2
			if t == array[ix] {
				return true
			} else if t < array[ix] {
				high = ix - 1
			} else {
				low = ix + 1
			}
			continue
		}
		break
	}
	return false
}
