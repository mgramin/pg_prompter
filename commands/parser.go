package commands

import (
	"fmt"
	"github.com/alecthomas/kong"
	"strings"
)

type Context struct {
	Debug bool
}

func Parse(cmd string) (*kong.Context, error) {
	parser, err := kong.New(&Cli, kong.UsageOnError())
	if err != nil {
		panic(err)
	}
	context, err := parser.Parse(strings.Fields(cmd))

	if err != nil {
		if parseError, ok := err.(*kong.ParseError); ok {
			context = parseError.Context
		} else {
			fmt.Println(err)
			return nil, nil
		}
	}

	return context, nil
}

func Run(context *kong.Context) error {
	err := context.Run(&Context{Debug: false})
	if err != nil {
		return err
	}
	return nil
}
