/*

*/
package sort

func getK(a []int) int {
	if len(a) == 0 {
		return 1
	}
	k := a[0]
	for _, v := range a {
		if v > k {
			k = v
		}
	}
	return k+1
}

func CountingSort(array []int) {
	l := len(array)
	k := getK(array)
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
