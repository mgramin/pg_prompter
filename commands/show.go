package commands

import (
	_ "embed"
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/olekukonko/tablewriter"
	"os"
	"pg_prompter/db"
	"strings"
)

//go:embed table_columns.sql
var tableColumns string

//go:embed all_objects.sql
var allObjects string

const ShowCommandName string = "show"

type ShowCmd struct {
	Pattern []string `arg:"" name:"pattern" help:"Object name pattern"`
}

type ShowCmdSuggester struct {
	suggestedObjects []prompt.Suggest
	next             Suggester
}

func (c ShowCmdSuggester) GetSuggestions(document prompt.Document) []prompt.Suggest {
	context, err := Parse(document.TextBeforeCursor())
	if err != nil {
	}
	if context.Path[1].Command.Name == ShowCommandName {
		return prompt.FilterFuzzy(c.suggestedObjects, document.GetWordBeforeCursor(), true)
	} else if c.next != nil {
		return c.next.GetSuggestions(document)
	} else {
		return []prompt.Suggest{}
	}
}

func (c ShowCmdSuggester) Refresh() {
}

func NewShowCmdSuggester(next Suggester) ShowCmdSuggester {
	var dbObjects []DbObject
	gorm, err := db.OpenGorm()
	CheckError(err)
	gorm.Raw(allObjects).Scan(&dbObjects)

	var s []prompt.Suggest
	for _, object := range dbObjects {
		s = append(s, prompt.Suggest{Text: object.Name, Description: object.Description})
	}

	return ShowCmdSuggester{s, next}
}

type TableDescription struct {
	Position string `caption:"#"`
	Column   string `caption:"Column"`
	Type     string `caption:"Type"`
	Length   string `caption:"Length"`
	Nullable string `caption:"Nullable"`
}

func (r *ShowCmd) Run(ctx *Context) error {
	gorm, err := db.OpenGorm()

	for _, p := range r.Pattern {
		split := strings.Split(p, ".")
		schemaName := split[0]
		objectName := split[1]
		CheckError(err)
		var result []TableDescription
		gorm.Raw(tableColumns, schemaName, objectName).Scan(&result)
		ShowTable(result, schemaName, objectName)
	}

	return nil
}

func ShowTable(data []TableDescription, schemaName string, objectName string) {
	fmt.Printf("%s.%s:\n", schemaName, objectName)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetBorder(false)
	table.SetHeader([]string{"Column", "Type", "Collation", "Nullable", "Default"})
	for _, v := range data {
		table.Append([]string{v.Column, v.Type, "", v.Nullable, ""})
	}
	table.Render()
	println()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
