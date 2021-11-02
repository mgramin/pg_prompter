package commands

import (
	"github.com/c-bata/go-prompt"
	"reflect"
)

var Cli struct {
	Describe    ShowCmd        `cmd:"" name:"show" help:"Describe detail information" aliases:"\d"`
	Drop        DropCmd        `cmd:"" name:"drop" help:"Drop DB object"`
	Connect     ConnectCmd     `cmd:"" name:"connect" help:"Connect to DB" aliases:"\c"`
	Gather      GatherCmd      `cmd:"" name:"gather" help:"Gather stat"`
	Vacuum      VacuumCmd      `cmd:"" name:"vacuum" help:"Vacuum"`
	Connections ConnectionsCmd `cmd:"" name:"connections" help:"List all predefined connections"`
	Exit        ExitCmd        `cmd:"" name:"exit" help:"Exit"`
}

type Suggester interface {
	GetSuggestions(document prompt.Document) []prompt.Suggest
	Refresh()
}

type DbObject struct {
	Name        string `gorm:"column:object_name"`
	Description string `gorm:"column:object_type"`
}

type CommandsSuggester struct {
	suggestedCommands []prompt.Suggest
	next              Suggester
}

func (c CommandsSuggester) GetSuggestions(document prompt.Document) []prompt.Suggest {
	context, err := Parse(document.TextBeforeCursor())
	if err != nil {
	}

	//if strings.HasSuffix(document.GetWordBeforeCursorWithSpace(), " ") {
	//	return []prompt.Suggest{}
	//}

	if len(context.Path) == 1 {
		return prompt.FilterFuzzy(c.suggestedCommands, document.GetWordBeforeCursor(), true)
	} else {
		return c.next.GetSuggestions(document)
	}
}

func (c CommandsSuggester) Refresh() {
}

func NewCommandsSuggester(next Suggester) CommandsSuggester {
	val := reflect.Indirect(reflect.ValueOf(&Cli))
	var s []prompt.Suggest
	for i := 0; i < val.Type().NumField(); i++ {
		s = append(s, prompt.Suggest{
			Text:        val.Type().Field(i).Tag.Get("name"),
			Description: val.Type().Field(i).Tag.Get("help")})
	}

	return CommandsSuggester{s, next}

}
