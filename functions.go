package testgolib

import "strings"

func Replace(s string) string {
	return strings.Replace(s, "Hello", "Goodbye", -1)
}
