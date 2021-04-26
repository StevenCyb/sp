package action

import (
	"fmt"
	"sp/model"
	"sp/utils"

	"github.com/mkideal/cli"
)

// Get action function
func Get(ctx *cli.Context) error {
	var err error
	var data *interface{}
	var result *interface{}
	argv := ctx.Argv().(*model.SelectT)

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

	result, err = utils.Get(data, argv.Query)
	if err != nil {
		return err
	}

	fmt.Println(*result)

	return nil
}
