package util

import (
	"regexp"
)

var usernameRegex = regexp.MustCompile(`^[\w\-.]{3,20}$`)

// IsValidUsername checks if the username is valid
func IsValidUsername(username string) bool {
	return usernameRegex.MatchString(username)
}
