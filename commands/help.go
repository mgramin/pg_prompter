package commands

import "fmt"

type HelpCmd struct {
}

func (r *HelpCmd) Run(ctx *Context) error {

	fmt.Println("Help me ...")

	return nil
}
