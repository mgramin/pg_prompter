package commands

// add pattern for all connection parameters

import (
	"database/sql"
	"fmt"
	"github.com/c-bata/go-prompt"
	"pg_prompter/db"
)

const ConnectCommandName string = "connect"

type ConnectCmd struct {
	Host     string `short:"H" long:"host" help:"Database server host or socket directory." default:""`
	Port     int    `short:"p" long:"port" help:"Database server port." default:""`
	Username string `short:"U" long:"username" help:"Database user name." default:""`
	Database string `short:"d" long:"dbname" help:"Database name to connect to." default:""`
}

func (r *ConnectCmd) Run(ctx *Context) error {
	var host string
	var port int
	var username string
	var database string
	var password string

	if r.Host != "" {
		host = r.Host
	} else {
		host = db.CurrentParams.Host
	}

	if r.Port != 0 {
		port = r.Port
	} else {
		port = db.CurrentParams.Port
	}

	if r.Username != "" {
		username = r.Username
	} else {
		username = db.CurrentParams.User
	}

	if r.Database != "" {
		database = r.Database
	} else {
		database = db.CurrentParams.Dbname
	}

	password = db.GetPassword(host, db.CurrentParams.User)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	dbconn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return err
	}

	err = dbconn.Ping()
	if err != nil {
		return err
	}

	db.CurrentParams.Host = host
	db.CurrentParams.Port = port
	db.CurrentParams.Dbname = database
	db.CurrentParams.User = username
	db.CurrentParams.Password = password

	return nil
}

type ConnectCmdSuggester struct {
	suggestedFlags []prompt.Suggest
	next           Suggester
}
