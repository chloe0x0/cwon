package main

import (
	"fmt"
	"os"
)

func main() {
	// shift the args to remove the program
	prompt := os.Args[1:]
	fmt.Println("prompt: ", prompt)
}