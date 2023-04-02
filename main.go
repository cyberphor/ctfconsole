package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	count     int
	printHelp bool
}

var helpString = fmt.Sprintf("")

func printHelp(w io.Writer) {
	fmt.Fprintf(w, helpString)
}

func validateArgs(c config) error {
	if !(c.count > 0) {
		return errors.New("Specify a number greater than 0.")
	}
	return nil
}

func parseArgs(args []string) (config, error) {
	var count int
	var err error
	c := config{}
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}
	if args[0] == "-h" || args[0] == "--help" {
		c.printHelp = true
		return c, nil
	}
	count, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.count = count
	return c, nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "What is your name? "
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("Enter a name.")
	}
	return name, nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s.\n", name)
	for i := 0; i < c.count; i++ {
		fmt.Fprintf(w, msg)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printHelp {
		printHelp(w)
		return nil
	}
	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printHelp(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printHelp(os.Stdout)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

/*
package main

import (
	"github.com/cyberphor/ctfconsole/models"
	"github.com/cyberphor/ctfconsole/views"
)

func main() {
	models.CreateTableForUsers()
	models.CreateTableForAdmins()
	models.CreateAdmin("admin", "password", "admin")
	views.Console()
}
*/
