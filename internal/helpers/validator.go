package helpers

import "regexp"

func IsUsernameValid(name string) (bool, string) {
	if len(name) < 5 || len(name) > 14 {
		return false, "Username must be between 5 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(name) {
		return false, "Username must contain only letters and numbers"
	}

	return true, ""
}

func IsNicknameValid(name string) (bool, string) {
	if len(name) < 8 || len(name) > 14 {
		return false, "Nickname must be between 8 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(name) {
		return false, "Nickname must contain only letters and numbers"
	}

	return true, ""
}

func IsPasswordValid(password string) (bool, string) {
	if len(password) < 8 || len(password) > 14 {
		return false, "Password must be between 8 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+$`).MatchString(password) {
		return false, "Password must contain letters, numbers and special characters"
	}

	return true, ""
}

func IsChannelNameValid(name string) (bool, string) {
	if len(name) < 2 || len(name) > 14 {
		return false, "Channel name must be between 2 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(name) {
		return false, "Channel name must contain only letters and numbers"
	}

	return true, ""
}
