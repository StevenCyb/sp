package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the validate command
var _ = Root.Register(&cli.Command{
	Name:    "validate",
	Aliases: []string{"val"},
	Desc:    "Validate provided structured data",
	Argv:    func() interface{} { return new(model.ValidateT) },
	Fn:      action.Validate,
})
