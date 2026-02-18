package helpers

import "regexp"

func IsNicknameValid(name string) (bool, string) {
	if len(name) < 8 || len(name) > 14 {
		return false, "Name must be between 8 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
		return false, "Name must contain only letters"
	}

	return true, ""
}
