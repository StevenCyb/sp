package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the merge command
var _ = Root.Register(&cli.Command{
	Name:    "merge",
	Aliases: []string{"mer"},
	Desc:    "Merge multiple strictured sources",
	Argv:    func() interface{} { return new(model.MergeT) },
	Fn:      action.Merge,
})
