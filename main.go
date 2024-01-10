package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
)

var (
	alt bool
	mu  sync.Mutex
)

func main() {
	// Read from standard input
	go readInput()

	// Read from standard error
	go readError()

	// Wait for both goroutines to finish
	select {}
}

func readInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		mu.Lock()
		alt = alternate(alt, scanner.Text())
		mu.Unlock()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
		os.Exit(1)
	}
}

func readError() {
	scanner := bufio.NewScanner(os.Stderr)

	for scanner.Scan() {
		mu.Lock()
		alt = alternate(alt, scanner.Text())
		mu.Unlock()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard error:", err)
		os.Exit(1)
	}
}

func alternate(alt bool, txt string) bool {
	if alt {
		color.Red(txt)
		alt = false
		return alt
	}
	color.Green(txt)
	alt = true
	return alt
}
