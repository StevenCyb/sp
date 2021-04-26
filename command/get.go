package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the get command
var _ = Root.Register(&cli.Command{
	Name: "get",
	Desc: "Get a specific element of strcutred data",
	Argv: func() interface{} { return new(model.SelectT) },
	Fn:   action.Get,
})
