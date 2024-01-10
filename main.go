package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var alt bool
	for scanner.Scan() {
		if alt {
			color.Red(scanner.Text())
			alt = false
			continue
		}
		color.Green(scanner.Text())
		alt = true
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
		os.Exit(1)
	}
}
