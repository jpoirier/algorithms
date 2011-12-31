/*
 Copyright (c) 2010 Joseph D Poirier
 Distributable under the terms of The New BSD License
 that can be found in the LICENSE file.
*/
package searching


func SequentialSearch(array []int, t int) bool {
	for i := len(array)-1; i >= 1; i-- {
		if array[i] == t {
			return true
		}
	}
	return false
}
