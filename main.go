package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "start" {
		fmt.Println("pass in 'start' as argument")
		return
	}

	bot := Bot{100, 100}

	scanner := bufio.NewScanner(os.Stdin)

	var gameLoop bool

	for {
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input. Try again later!:", err)
		}

		gameLoop = bot.converse(input)

		if !gameLoop {
			break
		}

	}
}
