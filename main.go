package main

import "github.com/matthiasharzer/helpme/commands"

func main() {
	err := commands.RootCommand.Execute()
	if err != nil {
		panic(err)
	}
}
