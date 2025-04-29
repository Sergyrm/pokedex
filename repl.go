package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	splittedSlice := strings.Fields(text)

	return splittedSlice
}