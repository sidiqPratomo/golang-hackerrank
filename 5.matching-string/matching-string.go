package matchingstring

import "fmt"

func MatchingStrings(stringList []string, queries []string) []int32 {
	// Write your code here
	freq := make(map[string]int32)
	results := make([]int32, len(queries))

	for _, str := range stringList {
		freq[str]++
	}
	fmt.Println(freq)
	for i, query := range queries {
		results[i] = freq[query]
	}
	return results
}