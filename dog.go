package godog

import (
	"strings"
)

func WhenGrownUp(s string) string {
	return "When the puppy grows up is says: " + strings.ToUpper(s)
}
