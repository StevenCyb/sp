package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the create command
var _ = Root.Register(&cli.Command{
	Name:    "create",
	Aliases: []string{"crt"},
	Desc:    "Create a new empty file",
	Argv:    func() interface{} { return new(model.CreateT) },
	Fn:      action.Create,
})
