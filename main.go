package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"github.com/olekukonko/tablewriter"
	"github.com/tg/pgpass"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

//go:embed data/all_objects.sql
var allObjects string

//go:embed data/table_columns.sql
var tableColumns string

//go:embed data/logo
var logo string

func completer1(d prompt.Document) []prompt.Suggest {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbParams.host, dbParams.port, dbParams.user, dbParams.password, dbParams.dbname)

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

	return prompt.FilterContains(s, d.GetWordBeforeCursor(), true)
}

func completer2(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "show", Description: "Show detail information"},
		{Text: "drop", Description: "Drop object"},
		{Text: "gather", Description: "Gather stat about object"},
		{Text: "vacuum", Description: "Vacuum"},
		{Text: "exit", Description: "Exit"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

var opts struct {
	Host     string `short:"H" long:"host" description:"database server host or socket directory" default:"localhost" value-name:"HOSTNAME"`
	Port     int    `short:"p" long:"port" description:"database server port" default:"5432" value-name:"PORT"`
	Username string `short:"U" long:"username" description:"database user name" default:"postgres" value-name:"USERNAME"`
	Database string `short:"d" long:"dbname" description:"database name to connect to" default:"postgres" value-name:"DBNAME"`
}

func main() {

	args := os.Args[1:]
	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		return
	}

	password, err := pgpass.Password(opts.Host, opts.Username)
	if err != nil {
		panic(err)
	}

	if password == "" {
		fmt.Printf("Password for user %v:\n", opts.Username)
		inputPassword, _ := terminal.ReadPassword(0)
		password = string(inputPassword)
	}

	dbParams = DbParams{opts.Host, opts.Port, opts.Username, opts.Database, password}

	fmt.Println(logo)

	fmt.Println("Please select command:")
	t := prompt.Input("> ", completer2, prompt.OptionCompletionOnDown())
	fmt.Println("You selected " + t)

	fmt.Println("Please select DB object:")
	currentObject := prompt.Input("> ", completer1, prompt.OptionCompletionOnDown(), prompt.OptionMaxSuggestion(12))
	fmt.Println("You selected " + currentObject)

	split := strings.Split(currentObject, ".")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbParams.host, dbParams.port, dbParams.user, dbParams.password, dbParams.dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	rows, err := db.Query(tableColumns, split[0], split[1])
	CheckError(err)

	defer rows.Close()

	var data [][]string

	i := 0
	for rows.Next() {
		a := []string{"", "", "", ""}

		err := rows.Scan(&a[0], &a[1], &a[2], &a[3])
		CheckError(err)

		data = append(data, a)

		i++
	}
	CheckError(rows.Err())

	ShowTable(data)

}

func ShowTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Column", "Type", "Nullable"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
