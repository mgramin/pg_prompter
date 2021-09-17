package commands

import (
	"fmt"
	"os"
)

type ExitCmd struct {
}

func (r *ExitCmd) Run(ctx *Context) error {
	fmt.Println("Bye bye ...")
	os.Exit(0)
	return nil
}
