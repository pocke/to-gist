package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ogier/pflag"
)

type CLI struct {
	Public      bool
	FileName    string
	FileContent string
}

func FlagParse(args []string) (*CLI, error) {
	fs := pflag.NewFlagSet(args[0], pflag.ContinueOnError)

	cli := &CLI{}

	fs.BoolVarP(&cli.Public, "private", "p", false, "make private")

	err := fs.Parse(args[1:])
	if err != nil {
		return nil, err
	}
	cli.Public = !cli.Public

	if len(fs.Args()) != 1 {
		return nil, fmt.Errorf("file name is required")
	}
	cli.FileName = fs.Arg(0)

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	cli.FileContent = string(content)

	return cli, nil
}
