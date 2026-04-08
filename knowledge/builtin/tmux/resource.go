package tmux

import (
	"github.com/matthiasharzer/helpme/knowledge/resource"
)

var Resource = resource.NestedResource{
	Items: []resource.Resource{
		resource.DescriptiveResource{
			Title: "tmux",
			Items: []resource.Item{
				PrefixInfoItem{BeforePrefix: "Your configured tmux prefix:"},
			},
		},
		resource.DescriptiveResource{
			Title: "tmux – session management",
			Items: []resource.Item{
				resource.ActionPairItem{
					Action:      "tmux new -s <session_name>",
					Description: "Create a new tmux session with the specified name",
				},
				resource.ActionPairItem{
					Action:      "tmux attach -t <session_name>",
					Description: "Attach to an existing tmux session",
				},
				resource.ActionPairItem{
					Action:      "tmux ls",
					Description: "List all tmux sessions",
				},
				resource.ActionPairItem{
					Action:      "tmux kill-session -t <session_name>",
					Description: "Kill the specified tmux session",
				},
				PrefixResource{
					Action:      "$",
					Description: "Rename session",
				},
			},
		},
		resource.DescriptiveResource{
			Title: "tmux key bindings – panes",
			Items: []resource.Item{
				PrefixResource{
					Action:      "\"",
					Description: "Split the current pane horizontally ↕",
				},
				PrefixResource{
					Action:      "%",
					Description: "Split the current pane vertically ↔",
				},
				PrefixResource{
					Action:      "o",
					Description: "Navigate to the next pane",
				},
				PrefixResource{
					Action:      ";",
					Description: "Navigate to the previous pane",
				},
				PrefixResource{
					Action:      "z",
					Description: "Toggle zoom for the current pane",
				},
				PrefixResource{
					Action:      "x",
					Description: "Kill the current pane",
				},
				PrefixResource{
					Action:      "Space",
					Description: "Toggle between pane layouts",
				},
				PrefixResource{
					Action:      "{",
					Description: "Swap the current with the left (or upper) pane",
				},
				PrefixResource{
					Action:      "}",
					Description: "Swap the current with right (or lower) pane",
				},
			},
		},
		resource.DescriptiveResource{
			Title: "tmux key bindings – windows",
			Items: []resource.Item{
				PrefixResource{
					Action:      "c",
					Description: "Create a new window",
				},
				PrefixResource{
					Action:      "d",
					Description: "Detach from session",
				},
				PrefixResource{
					Action:      "w",
					Description: "List all tmux windows",
				},
				PrefixResource{
					Action:      "&",
					Description: "Kill the current window",
				},
				PrefixResource{
					Action:      "n",
					Description: "Navigate to the next window",
				},
				PrefixResource{
					Action:      "p",
					Description: "Navigate to the previous window",
				},
				PrefixResource{
					Action:      ",",
					Description: "Rename the current window",
				},
				PrefixResource{
					Action:      "0...9",
					Description: "Switch to window",
				},
			},
		},
		resource.DescriptiveResource{
			Title: "tmux key bindings – utility",
			Items: []resource.Item{
				PrefixResource{
					Action:      "t",
					Description: "Show the current time",
				},
				PrefixResource{
					Action:      "?",
					Description: "Show all key bindings",
				},
			},
		},
		resource.DescriptiveResource{
			Title: "tmux commands",
			Items: []resource.Item{
				PrefixResource{
					Action:      ":  or  tmux <command>",
					Description: "Run a tmux command",
				},
				resource.ActionPairItem{
					Action:      ":set status (on|off)",
					Description: "Toggle status bar",
				},
				resource.ActionPairItem{
					Action:      ":set status-position (top|bottom)",
					Description: "Set the status bar position",
				},
				resource.ActionPairItem{
					Action:      `:set status-style "bg=<color>,fg=<color>"`,
					Description: "Set status bar colors",
				},
			},
		},
	},
}
