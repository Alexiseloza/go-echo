package validator

import (
	"regexp"
	"unicode"
)

var Regex_email = regexp.MustCompile("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")

func ValidatePassword(s string) bool {
	var (
		hasMinLen = false
		hasUpper  = false
		hasLower  = false
		hasNumber = false
	)
	if len(s) >= 6 && len(s) <= 20 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
			//case unicode.IsPunct(char) || unicode.IsSymbol(char):
			//hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber
}
