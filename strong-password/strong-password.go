package strongpassword

import (
	"unicode"
	"strings"
)

func MinimumNumber(n int32, password string) int32 {
    // Return the minimum number of characters to make the password strong
    var (
        hasLower bool
        hasUpper bool
        hasDigit bool
        hasSpecial bool
    )

    specialChars := "!@#$%^&*()-+"
    
    for _, char := range password {
        switch {
        case unicode.IsLower(char):
            hasLower = true
        case unicode.IsUpper(char):
            hasUpper = true
        case unicode.IsDigit(char):
            hasDigit = true
        case strings.ContainsRune(specialChars, char):
            hasSpecial = true
        }
    }
    requiredCategories := 0
    if !hasLower {
        requiredCategories++
    }
    if !hasUpper {
        requiredCategories++
    }
    if !hasDigit {
        requiredCategories++
    }
    if !hasSpecial {
        requiredCategories++
    }
    requiredLength := int32(0)
    if n < 6 {
        requiredLength = 6 - n
    }
    if requiredLength > int32(requiredCategories) {
        return requiredLength
    }
    return int32(requiredCategories)
}