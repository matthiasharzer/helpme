package commands

import (
	"github.com/matthiasharzer/helpme/knowledge"
	"github.com/spf13/cobra"
)

var version = "unknown"

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
			cmd.Println("helpme version", version)
			return nil
		}

		if len(args) != 1 {
			cmd.Println("Please specify a command.")
			return cmd.Help()
		}

		command := args[0]

		resources := knowledge.LoadResources()

		if resource, exists := resources[command]; exists {
			cmd.Println(resource.Content())
		} else {
			cmd.Println("Command not found.")
			return cmd.Help()
		}

		return cmd.Help()
	},
}
