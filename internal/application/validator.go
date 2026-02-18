package application

import "regexp"

//TODO check file naming, also maybe move these to separate helpers folder?

func isUsernameValid(name string) (bool, string) {
	if len(name) < 8 || len(name) > 14 {
		return false, "Username must be between 8 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(name) {
		return false, "Username must contain only letters and numbers"
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

func isChannelNameValid(name string) (bool, string) {
	if len(name) < 2 || len(name) > 14 {
		return false, "Channel name must be between 2 and 14 characters long"
	}

	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(name) {
		return false, "Channel name must contain only letters and numbers"
	}

	return true, ""
}
