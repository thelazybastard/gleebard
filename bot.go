package main

import (
	"slices"
	"strings"
)

type Bot struct {
	Health    int
	Annoyance int
}

func (bot *Bot) converse(input string) bool {
	if strings.Contains(input, "end") {
		return false
	}

	mapKey := check_input(input)

	return true
}

func check_input(input string) string {
	for key, value := range script {
		if slices.Contains(value, input) {
			return key
		}
	}

	return "UNKNOWN"
}
