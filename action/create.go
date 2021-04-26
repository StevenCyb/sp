package action

import (
	"sp/model"
	"sp/utils"

	"github.com/mkideal/cli"
)

// Create action function
func Create(ctx *cli.Context) error {
	var err error
	argv := ctx.Argv().(*model.CreateT)

	newEmpty := utils.CreateEmptyMap()
	for _, file := range argv.FileOutput {
		err = utils.DataToFile(newEmpty, argv.OutputFormat, file, false)
		if err != nil {
			return err
		}
	}

	return nil
}
