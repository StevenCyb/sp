package action

import (
	"fmt"
	"sp/model"
	"sp/utils"

	"github.com/mkideal/cli"
)

// Put action function
func Put(ctx *cli.Context) error {
	var err error
	var data *interface{}
	argv := ctx.Argv().(*model.PutT)

	if argv.StdInput != "" {
		data, err = utils.NewDataFromString(argv.StdInput, argv.InputFormat)
	} else if argv.FileInput != "" {
		data, err = utils.NewDataFromFile(argv.FileInput, argv.InputFormat)
	} else {
		return fmt.Errorf("no structured input data provided")
	}

	if err != nil {
		return err
	}

	err = utils.Put(data, argv.Query, utils.StringCaster(argv.Item))
	if err != nil {
		return err
	}

	if argv.FileOutput != "" {
		err = utils.DataToFile(data, argv.OutputFormat, argv.FileOutput, argv.PrettyOutput)
	} else if argv.StdOutput {
		err = utils.DataToStd(data, argv.OutputFormat, argv.PrettyOutput)
	}

	return err
}
