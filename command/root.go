package command

import (
	"os"
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Root - register the root command
var Root = &cli.Command{
	Name: os.Args[0],
	Desc: "SP",
	Text: `Process structured data from command line`,
	Argv: func() interface{} { return new(model.RootT) },
	Fn:   action.Root,
}
