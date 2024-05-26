package twocharacter

func Alternate(s string) int32 {
	// Write your code here
	uniqueChars := make(map[rune]bool)
	for _, char := range s {
		uniqueChars[char] = true
	}

	chars := []rune{}
	for char := range uniqueChars {
		chars = append(chars, char)
	}

	maxLen := int32(0)

	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			char1, char2 := chars[i], chars[j]
			prevChar := rune(0)
			valid := true
			currentLen := int32(0)

			for _, char := range s {
				if char == char1 || char == char2 {
					if char == prevChar {
						valid = false
						break
					}
					prevChar = char
					currentLen++
				}
			}

			if valid && currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}
	return maxLen
}