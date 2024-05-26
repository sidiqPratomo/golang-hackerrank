package birthdaycakecandle

func BirthdayCakeCandles(candles []int32) int32 {
	var maxHeight int32 = 0
	var count int32 = 0

	for _, height := range candles{
		if height > maxHeight{
			maxHeight = height
			count = 1
		} else if height == maxHeight {
			count++
		}
	}
	return count
}