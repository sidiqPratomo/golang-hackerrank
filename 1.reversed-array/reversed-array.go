package reversedarray

func Reversed(a []int32) []int32{
	length :=len(a)

	reversed := make([]int32,length)

	for i:=0;i<length;i++{
		reversed[i] = a[length-1-i]
	}
	return reversed
}