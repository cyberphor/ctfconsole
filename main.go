package main

import (
	"fmt"
	"os"

	"https://github.com/cyberphor/ctfconsole/cmd"
)

func main() {
	c, err := cmd.parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	fmt.Println(c)
}
