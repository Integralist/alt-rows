package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

var (
	g = ansi.ColorFunc("green+bh")
	r = ansi.ColorFunc("red+bh")
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var alt bool
	for scanner.Scan() {
		if alt {
			fmt.Println(r(scanner.Text()))
			alt = false
			continue
		}
		fmt.Println(g(scanner.Text()))
		alt = true
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
		os.Exit(1)
	}
}
