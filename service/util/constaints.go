package util

import (
	"regexp"
)

var usernameRegex = regexp.MustCompile(`^[\w\-.]{3,20}$`)
var displayNameRegex = regexp.MustCompile(`^.{3,32}$`)

// IsValidUsername checks if the username is valid
func IsValidUsername(username string) bool {
	return usernameRegex.MatchString(username)
}

// IsValidDisplayName checks if the display name is valid
func IsValidDisplayName(displayName string) bool {
	return displayNameRegex.MatchString(displayName)
}
