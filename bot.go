package main

type Bot struct {
	Health    int
	Annoyance int
}

func (bot *Bot) converse(input string) bool {
	if input == "end" {
		return false
	}
	return true
}
