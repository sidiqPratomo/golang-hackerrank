package timeconversion

import "strconv"

func TimeConversion(s string) string {
	// Write your code here
	period := s[len(s)-2:]
	hour := s[:2]
	minute := s[3:5]
	second := s[6:8]

	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		return "error"
	}

	if period == "AM" {
		if hourInt == 12 {
			hourInt = 0
		}
	} else if period == "PM" {
		if hourInt != 12 {
			hourInt += 12
		}
	}
	newHour := strconv.Itoa(hourInt)
	if hourInt < 10 {
		newHour = "0" + newHour
	}
	return newHour + ":" + minute + ":" + second
}