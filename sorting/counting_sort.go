/*

*/
package sort


func getK(a []uint) int {
	if len(a) == 0 {
		return 1
	}
	k := a[0]
	for _, v := range a {
		if v > k {
			k = v
		}
	}
	return int(k+1)
}

func CountingSort(array []uint) {
	l := len(array)
	k := getK(array)
	b := make([]uint, k)	// array of counts
	for i:= 0; i < l; i++ {
		b[array[i]] += uint(1)
	}
	for i, j := 0, 0; i < k; i++ {
		for {
			if b[i] > 0 {
				array[j] = uint(i)
				j += 1
				b[i] -= uint(1)
				continue
			}
			break
		}
	}
}
