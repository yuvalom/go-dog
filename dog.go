package godog

import (
	"strings"
)

func WhenGrownUp(s string) string {
	return "When the puppy grows up is says: " + strings.ToUpper(s)
}

func From1_0() string {
	return "I am from version v1.0.0"
}
