package commands

import (
	"github.com/alecthomas/kong"
	"strings"
)

type Context struct {
	Debug bool
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Show ShowCmd `cmd help:"Show."`
	Help HelpCmd `cmd help:"Help."`
	Exit ExitCmd `cmd help:"Exit."`
}

func Parse(cmd string) {
	parser, err := kong.New(&cli)
	if err != nil {
		panic(err)
	}
	ctx, err := parser.Parse(strings.Fields(cmd))
	parser.FatalIfErrorf(err)

	err = ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)

}
