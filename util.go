package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CustomOptions(prefix string) []prompt.Option {
	options := []prompt.Option{
		prompt.OptionTitle("pg-prompter: seek and destroy any DB object"),
		prompt.OptionPrefix(prefix),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionSuggestionBGColor(prompt.Purple),
		prompt.OptionDescriptionBGColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.Blue),
		prompt.OptionSelectedDescriptionBGColor(prompt.Purple),
		prompt.OptionSelectedSuggestionTextColor(prompt.Red),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	}
	return options
}
