package action

import (
	"fmt"
	"io/ioutil"
	"sp/model"
	"sp/utils"

	"github.com/mkideal/cli"
)

// Validate action function
func Validate(ctx *cli.Context) error {
	var isValid bool
	argv := ctx.Argv().(*model.ValidateT)

	if argv.StdInput != "" {
		isValid = utils.StringIsValidStructure(argv.StdInput, argv.InputFormat)
	} else if argv.FileInput != "" {
		data, err := ioutil.ReadFile(argv.FileInput)
		if err != nil {
			return err
		}

		isValid = utils.ByteIsValidStructure(data, argv.InputFormat)
	} else {
		return fmt.Errorf("no structured input data provided")
	}

	fmt.Println(isValid)

	return nil
}
