package action

import (
	"fmt"
	"sp/model"
	"sp/utils"

	"github.com/mkideal/cli"
)

// Difference action function
func Difference(ctx *cli.Context) error {
	var dataSlice []*interface{}
	argv := ctx.Argv().(*model.DifferenceT)

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
	if dataCount != 2 {
		return fmt.Errorf("you need to provide exact two inputs")
	}

	result, err := utils.Difference(dataSlice[0], dataSlice[1], argv.OutputFormat)

	if err == nil {
		fmt.Println(result)
	}

	return err
}
