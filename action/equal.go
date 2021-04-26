package action

import (
	"fmt"
	"sp/model"
	"sp/utils"

	"github.com/mkideal/cli"
)

// Equal action function
func Equal(ctx *cli.Context) error {
	var dataSlice []*interface{}
	argv := ctx.Argv().(*model.EqualT)

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
		if !utils.Equal(dataSlice[i], dataSlice[i+1]) {
			fmt.Println(false)
			return nil
		}
	}

	fmt.Println(true)

	return nil
}
