package main

import (
	"flag"
	"fmt"
)

func main() {
	var c = flag.String("c", "", "command")
	flag.Parse()
	fmt.Println(*c)
}
