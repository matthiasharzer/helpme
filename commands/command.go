package commands

import (
	"github.com/matthiasharzer/helpme/knowledge"
	"github.com/spf13/cobra"
)

var version string

var showVersion bool

func init() {
	RootCommand.Flags().BoolVarP(&showVersion, "version", "v", false, "Show version information")
}

var RootCommand = &cobra.Command{
	Use: "helpme [command]",
	Long: `helpme is a command-line tool that provides quick access to various resources and information.
It offers a range of commands to assist users in finding answers, accessing documentation, and more.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			println("helpme version", version)
			return nil
		}

		if len(args) != 1 {
			println("Please specify a command. Use 'helpme --help' for more information.")
			return nil
		}

		command := args[0]

		resources := knowledge.LoadResources()

		if resource, exists := resources[command]; exists {
			println(resource.Content())
		} else {
			println("Command not found. Use 'helpme --help' for more information.")
		}

		return nil
	},
}
