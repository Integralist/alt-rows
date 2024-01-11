package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/mgutz/ansi"
)

var (
	cg = ansi.ColorFunc("green+bh")
	cr = ansi.ColorFunc("red+bh")
)

func main() {
	var r *regexp.Regexp
	pattern := patternFlag()
	if pattern != "" {
		r = regexp.MustCompile(pattern)
	}

	if err := scan(r); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
		os.Exit(1)
	}
}

func patternFlag() string {
	var pattern string
	flag.StringVar(&pattern, "pattern", "", "specify regex pattern for lines to apply alternating rows to (defaults to all input)")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: alt-rows [-pattern=<regex>]")
		flag.PrintDefaults()
	}
	flag.Parse()
	return pattern
}

func scan(r *regexp.Regexp) error {
	scanner := bufio.NewScanner(os.Stdin)

	var alt bool
	for scanner.Scan() {
		if r == nil || r.Match(scanner.Bytes()) {
			if alt {
				fmt.Println(cr(scanner.Text()))
				alt = false
				continue
			}
			fmt.Println(cg(scanner.Text()))
			alt = true
		}
	}

	return scanner.Err()
}
