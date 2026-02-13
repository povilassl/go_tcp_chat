package hub

import "regexp"

func isNameValid(name string) (bool, string) {
	if len(name) < 8 || len(name) > 14 {
		return false, "Name must be between 8 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(name) {
		return false, "Name must contain only letters and numbers"
	}

	return true, ""
}

func isPasswordValid(password string) (bool, string) {
	if len(password) < 8 || len(password) > 14 {
		return false, "Password must be between 1 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+$`).MatchString(password) {
		return false, "Password must contain letters, numbers and special characters"
	}

	return true, ""
}
