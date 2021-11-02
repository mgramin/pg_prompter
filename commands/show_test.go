package commands

import (
	"github.com/c-bata/go-prompt"
	"testing"
)

func TestGetSuggestions(t *testing.T) {
	suggester := ShowCmdSuggester{nil, nil}

	buffer := prompt.NewBuffer()
	buffer.InsertText("describe table ", false, true)
	document := buffer.Document()

	//println(document.DisplayCursorPosition())

	suggester.GetSuggestions(*document)
}
