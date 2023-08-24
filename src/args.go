package cmd

import (
	"flag"
	"io"
)

type config struct {
	filePath string
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("ctfconsole", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.filePath, "c", ".", "Help message goes here")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	return c, nil
}
