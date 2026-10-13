package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "start" {
		fmt.Println("pass in 'start' as argument")
		return
	}

}
