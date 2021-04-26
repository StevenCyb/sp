package action

import (
	"sp/model"

	"github.com/mkideal/cli"
)

// version of this application
const version = "v0.0.1"

// Root action function
func Root(ctx *cli.Context) error {
	argv := ctx.Argv().(*model.RootT)

	// Print version if flag set
	if argv.Version {
		ctx.String(version + "\n")
		return nil
	}

	// Print command list if flag set
	if argv.List {
		ctx.String(ctx.Command().ChildrenDescriptions(" ", "  =>  "))
		return nil
	}

	// Print info if no arguments
	ctx.String("try `%s --help for more information'\n", ctx.Path())
	return nil
}
