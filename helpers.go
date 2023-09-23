package main

import (
	"fmt"
	"strings"
)

// SplitLastString returns a string representation of the last string input after last slash
// and also returns a string without the extension
func SplitLastString(text string) (string, error) {
	// Split the URL by "/"
	parts := strings.Split(text, "/")

	// Find the last part, which should be "grass.gif"
	lastPart := parts[len(parts)-1]

	// Split the last part by "."
	typeParts := strings.Split(lastPart, ".")

	// Check if there are at least two parts (e.g., "grass.gif")
	if len(typeParts) < 2 {
		return "", fmt.Errorf("invalid URL format")
	}

	// The type is the first part before the "."
	result := typeParts[0]

	return result, nil
}
