package main

import (
	_ "embed"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/c-bata/go-prompt"
	_ "github.com/lib/pq"
	"pg_prompter/commands"
	"pg_prompter/db"
)

func complete(d prompt.Document) []prompt.Suggest {
	return suggesters.GetSuggestions(d)
}

func exec(cmd string) {
	context, err := commands.Parse(cmd)
	if err != nil {
		fmt.Println(err)
	}

	err = commands.Run(context)
	if err != nil {
		fmt.Println(err)
	}

	suggesters.Refresh()
}

type initConnectParam struct {
	Host     string `short:"H" long:"host" help:"Database server host or socket directory." default:"localhost"`
	Port     int    `short:"p" long:"port" help:"Database server port." default:"5432"`
	Username string `short:"U" long:"username" help:"Database user name." default:"postgres"`
	Database string `short:"d" long:"dbname" help:"Database name to connect to." default:"postgres"`
}

var conOpts initConnectParam

var suggesters commands.Suggester

func main() {
	kong.Parse(&conOpts)

	password := db.GetPassword(conOpts.Host, conOpts.Username)
	db.CurrentParams = db.Params{Host: conOpts.Host, Port: conOpts.Port, User: conOpts.Username, Dbname: conOpts.Database, Password: password}

	suggesters = commands.NewCommandsSuggester(
		commands.NewShowCmdSuggester(nil))

	p := prompt.New(
		exec,
		complete,
		CustomOptions()...,
	)
	p.Run()

}
