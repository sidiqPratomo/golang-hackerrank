package minimaxarray

import (
	"fmt"
	"sort"
)

func MiniMaxSum(arr []int32) {
	// Write your code here
	intArr := make([]int, len(arr))
	for i, v := range arr {
		intArr[i] = int(v)
	}
	sort.Ints(intArr)
	minSum := 0
	for i := 0; i<4 ; i++ {
		minSum += intArr[i]
	}
	maxSum := 0
	for i := len(intArr) - 4; i < len(intArr); i++ {
		maxSum += intArr[i]
	}
	fmt.Println(minSum, maxSum)
}