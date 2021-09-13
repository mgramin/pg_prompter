package commands

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"pg_prompter/db"
	"strings"
)

//go:embed table_columns.sql
var tableColumns string

type ShowCmd struct {
	Force     bool `help:"Force removal."`
	Recursive bool `help:"Recursively remove files."`

	Paths string `arg name:"path" help:"Paths to remove."`
}

func (r *ShowCmd) Run(ctx *Context) error {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.CurrentParams.Host, db.CurrentParams.Port, db.CurrentParams.User, db.CurrentParams.Password, db.CurrentParams.Dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	split := strings.Split(r.Paths, ".")
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

	return nil
}

func ShowTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Column", "Type", "Nullable"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
