package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the delete command
var _ = Root.Register(&cli.Command{
	Name:    "delete",
	Aliases: []string{"del"},
	Desc:    "Delete an element of strcutred data",
	Argv:    func() interface{} { return new(model.DeleteT) },
	Fn:      action.Delete,
})
