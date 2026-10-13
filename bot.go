package main

import (
	"math/rand/v2"
	"slices"
)

type Bot struct {
	Health    int
	Annoyance int
}

func (bot *Bot) converse(input string) (string, bool) {
	dialogueSet := script[check_input(input)]

	randomChoice := rand.IntN(len(dialogueSet))
	response := dialogueSet[randomChoice]

	return response, true
}

func check_input(input string) string {
	for key, value := range script {
		if slices.Contains(value, input) {
			return key
		}
	}

	return "UNKNOWN"
}
