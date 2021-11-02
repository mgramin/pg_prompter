package commands

import (
	_ "embed"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"pg_prompter/db"
)

//go:embed databases.sql
var databases string

type ConnectionsCmd struct {
}

type DatabaseInfo struct {
	Name       string `caption:"Name"`
	Owner      string `caption:"Owner"`
	Encoding   string `caption:"Encoding"`
	Collate    string `caption:"Collate"`
	Ctype      string `caption:"Ctype"`
	Privileges string `caption:"Access privileges"`
}

func (r *ConnectionsCmd) Run(ctx *Context) error {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.CurrentParams.Host, db.CurrentParams.Port, db.CurrentParams.User, db.CurrentParams.Password, db.CurrentParams.Dbname)
	var result []DatabaseInfo
	db1, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	CheckError(err)
	db1.Raw(databases).Scan(&result)
	ShowDatabases(result)
	return nil
}

func ShowDatabases(data []DatabaseInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetBorder(false)
	table.SetHeader([]string{"Name", "Owner", "Encoding", "Collate", "Ctype", "Access privileges"})
	for _, v := range data {
		table.Append([]string{v.Name, v.Owner, v.Encoding, v.Collate, v.Ctype, v.Privileges})
	}
	table.Render()
}
