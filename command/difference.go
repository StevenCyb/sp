package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the difference command
var _ = Root.Register(&cli.Command{
	Name:    "difference",
	Aliases: []string{"diff"},
	Desc:    "Create structured difference of multiple strictured sources",
	Argv:    func() interface{} { return new(model.DifferenceT) },
	Fn:      action.Difference,
})
