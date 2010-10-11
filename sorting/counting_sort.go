/*

*/
package sort


func CountingSort(array []int, k int) {
	l := len(array)
	b := make([]int, k)	// array of counts
	for i:= 0; i < l; i++ {
		b[array[i]] += 1
	}
	for i, j := 0, 0; i < k; i++ {
		for {
			if b[i] > 0 {
				array[j] = i
				j += 1
				b[i] -= 1
				continue
			}
			break
		}
	}
}
