package rotateleft

func RotateLeft(d int32, arr []int32) []int32 {
    // Write your code here
    length :=len(arr)
    results := make([]int32, length)
    for i:=0; i<length; i++{
        newIndex := (i + int(length) - int(d)) % length
        results[newIndex] = arr[i]
    }
    return results
}