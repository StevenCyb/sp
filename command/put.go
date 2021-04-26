package command

import (
	"sp/action"
	"sp/model"

	"github.com/mkideal/cli"
)

// Register the put command
var _ = Root.Register(&cli.Command{
	Name: "put",
	Desc: "Put an item on specific position",
	Argv: func() interface{} { return new(model.PutT) },
	Fn:   action.Put,
})
