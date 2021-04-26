package main

import (
	"fmt"
	"os"
	"sp/command"

	"github.com/mkideal/cli"
)

func main() {
	cli.SetUsageStyle(cli.ManualStyle)
	if err := command.Root.RunWith(os.Args[1:], os.Stderr, nil); err != nil {
		fmt.Printf("%v\n", err)
	}
}
