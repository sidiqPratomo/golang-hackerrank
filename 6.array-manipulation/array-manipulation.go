package arraymanipulation

import "fmt"

func ArrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here
	arr := make([]int64, n)
	for _, query := range queries {
		a := query[0] - 1
		b := query[1] - 1
		k := int64(query[2])

		arr[a] += k
		if b+1 < n {
			arr[b+1] -= k
		}
		fmt.Println(arr)
	}
	maxVal := int64(0)
	curVal := int64(0)
	for _, val := range arr {
		fmt.Println(val)
		curVal += val
		if curVal > maxVal {
			maxVal = curVal
		}
	}
	return maxVal
}