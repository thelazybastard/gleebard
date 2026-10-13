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
	mapKey := check_input(input)
	dialogueSet := script[mapKey]

	randomChoice := rand.IntN(len(dialogueSet))
	response := dialogueSet[randomChoice]
	continueGame := true

	switch mapKey {
	case "USER_BLOODTHIRSTY":
		bot.Health -= 10
	case "USER_HEAL":
		bot.Health += 10
	case "MINIGAME_WIN", "SOLVED_RIDDLE":
		bot.Annoyance += 10
	case "COMPLIMENT":
		bot.Annoyance -= 10
	}

	if mapKey == "END" || bot.Health == 0 || bot.Annoyance == 0 {
		continueGame = false
	}

	return response, continueGame
}

func check_input(input string) string {
	for key, value := range script {
		if slices.Contains(value, input) {
			return key
		}
	}

	return "UNKNOWN"
}
