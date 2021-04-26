package action

import (
	"fmt"
	"sp/model"
	"sp/utils"

	"github.com/mkideal/cli"
)

// Merge action function
func Merge(ctx *cli.Context) error {
	var err error
	var dataSlice []*interface{}
	argv := ctx.Argv().(*model.MergeT)

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

	for i := 0; i < dataCount-1; i++ {
		utils.Merge(dataSlice[i], dataSlice[i+1], argv.AppendArray)
	}

	if argv.FileOutput != "" {
		err = utils.DataToFile(dataSlice[dataCount-1], argv.OutputFormat, argv.FileOutput, argv.PrettyOutput)
	} else if argv.StdOutput {
		err = utils.DataToStd(dataSlice[dataCount-1], argv.OutputFormat, argv.PrettyOutput)
	}

	return err
}
