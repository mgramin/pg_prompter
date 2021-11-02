package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"pg_prompter/db"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CustomOptions() []prompt.Option {
	options := []prompt.Option{
		prompt.OptionTitle("pg-prompter: seek and destroy any DB object"),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionSuggestionBGColor(prompt.Purple),
		prompt.OptionDescriptionBGColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.Blue),
		prompt.OptionSelectedDescriptionBGColor(prompt.Purple),
		prompt.OptionSelectedSuggestionTextColor(prompt.Red),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
		prompt.OptionLivePrefix(changeLivePrefix),
	}
	return options
}

func changeLivePrefix() (string, bool) {
	return fmt.Sprintf("%s@%s:%d>", db.CurrentParams.Dbname, db.CurrentParams.Host, db.CurrentParams.Port), true
}
