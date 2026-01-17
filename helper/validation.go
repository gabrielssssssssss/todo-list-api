package helper

import (
	"regexp"
)

func IsEmailValid(email string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	return re.MatchString(email)
}

func IsStrongerPassword(password string) bool {
	re := regexp.MustCompile(`^[A-Za-z\d!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?]{12,}$`)
	return re.MatchString(password)
}
