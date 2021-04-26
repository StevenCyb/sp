package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the equal command
var _ = Root.Register(&cli.Command{
	Name:    "equal",
	Aliases: []string{"eq"},
	Desc:    "Check if given structured data is equal",
	Argv:    func() interface{} { return new(model.EqualT) },
	Fn:      action.Equal,
})
