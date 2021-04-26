package action

import (
	"fmt"
	"sp/model"
	"sp/utils"

	"github.com/mkideal/cli"
)

// Concatenate action function
func Concatenate(ctx *cli.Context) error {
	var dataSlice []*interface{}
	argv := ctx.Argv().(*model.ConcatenateT)

	for _, filePath := range argv.FileInput {
		data, err := utils.NewDataFromFile(filePath, argv.InputFormat)
		if err != nil {
			return err
		}

		dataSlice = append(dataSlice, data)
	}

	for _, stdin := range argv.StdInput {
		data, err := utils.NewDataFromString(stdin, argv.InputFormat)
		if err != nil {
			return err
		}

		dataSlice = append(dataSlice, data)
	}

	dataCount := len(dataSlice)
	if dataCount <= 1 {
		return fmt.Errorf("you need to provide at least two inputs")
	}

	for _, data := range dataSlice {
		utils.DataToStd(data, argv.OutputFormat, argv.PrettyOutput)
	}

	return nil
}
