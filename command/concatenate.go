package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the concatenate command
var _ = Root.Register(&cli.Command{
	Name:    "concatenate",
	Aliases: []string{"cat"},
	Desc:    "Concatenate data",
	Argv:    func() interface{} { return new(model.ConcatenateT) },
	Fn:      action.Concatenate,
})
