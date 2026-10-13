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

	bot := Bot{100, 0}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input. Try again later!:", err)
		}

		resp, gameLoop := bot.converse(input)

		if !gameLoop {
			break
		}

		fmt.Println(resp)
	}
}
