package main

import (
	"fmt"
	"os"

	"github.com/matthiasharzer/helpme/commands"
)

func main() {
	err := commands.RootCommand.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
