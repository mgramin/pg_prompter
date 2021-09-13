package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/c-bata/go-prompt"
	_ "github.com/lib/pq"
	"github.com/tg/pgpass"
	"golang.org/x/crypto/ssh/terminal"
	"pg_prompter/commands"
	"pg_prompter/db"
)

//go:embed data/all_objects.sql
var allObjects string

//go:embed data/logo
var logo string

func completer1(d prompt.Document) []prompt.Suggest {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.CurrentParams.Host, db.CurrentParams.Port, db.CurrentParams.User, db.CurrentParams.Password, db.CurrentParams.Dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	var s []prompt.Suggest

	rows, err := db.Query(allObjects)
	CheckError(err)

	for rows.Next() {
		var suggest prompt.Suggest
		err := rows.Scan(&suggest.Text, &suggest.Description)
		CheckError(err)
		s = append(s, suggest)
	}
	err = rows.Close()
	CheckError(err)

	return prompt.FilterFuzzy(s, d.GetWordBeforeCursor(), true)
}

func completer2(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "show", Description: "Show detail information"},
		{Text: "drop", Description: "Drop object"},
		{Text: "gather", Description: "Gather stat about object"},
		{Text: "vacuum", Description: "Vacuum"},
		{Text: "connect", Description: "Connect"},
		{Text: "help", Description: "Help"},
		{Text: "exit", Description: "Exit"},
	}

	if d.FindStartOfPreviousWord() == 0 {
		return prompt.FilterFuzzy(s, d.GetWordBeforeCursor(), true)
	} else {
		return completer1(d)
	}

}

func exec(cmd string) {
	commands.Parse(cmd)
}

func main() {

	kong.Parse(&commands.ConOpts)

	password, err := pgpass.Password(commands.ConOpts.Host, commands.ConOpts.Username)
	if err != nil {
		panic(err)
	}

	if password == "" {
		fmt.Printf("Password for user %v:\n", commands.ConOpts.Username)
		inputPassword, _ := terminal.ReadPassword(0)
		password = string(inputPassword)
	}

	db.CurrentParams = db.Params{Host: commands.ConOpts.Host, Port: commands.ConOpts.Port, User: commands.ConOpts.Username, Dbname: commands.ConOpts.Database, Password: password}

	//fmt.Println(logo)

	prefix := fmt.Sprintf("%s@%s:%d>", db.CurrentParams.Dbname, db.CurrentParams.Host, db.CurrentParams.Port)
	p := prompt.New(
		exec,
		completer2,
		CustomOptions(prefix)...,
	)
	p.Run()

}
