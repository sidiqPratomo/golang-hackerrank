package hourglasssum

import "math"

func HourglassSum(arr [][]int32) int32 {
	// Write your code here
	maxSum := math.MinInt32
	maxSumI32 := int32(maxSum)
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] +
				arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if sum >= maxSumI32 {
				maxSumI32 = sum
			}
		}
	}
	return maxSumI32
}