package knowledge

func LoadResources() map[string]Resource {
	return map[string]Resource{
		"tmux": nestedResource{
			items: []Resource{
				descriptiveResource{
					title: "tmux – session management",
					items: []resourceItem{
						actionPair{
							action:      "tmux new -s <session_name>",
							description: "Create a new tmux session with the specified name",
						},
						actionPair{
							action:      "tmux attach -t <session_name>",
							description: "Attach to an existing tmux session",
						},
						actionPair{
							action:      "tmux ls",
							description: "List all tmux sessions",
						},
						actionPair{
							action:      "tmux kill-session -t <session_name>",
							description: "Kill the specified tmux session",
						},
						actionPair{
							action:      "Ctrl + b, $",
							description: "Rename session",
						},
					},
				},
				descriptiveResource{
					title: "tmux key bindings – panes",
					items: []resourceItem{
						actionPair{
							action:      "Ctrl + b, \"",
							description: "Split the current pane horizontally ↕",
						},
						actionPair{
							action:      "Ctrl + b, %",
							description: "Split the current pane vertically ↔",
						},
						actionPair{
							action:      "Ctrl + b, o",
							description: "Navigate to the next pane",
						},
						actionPair{
							action:      "Ctrl + b, ;",
							description: "Navigate to the previous pane",
						},
						actionPair{
							action:      "Ctrl + b, z",
							description: "Toggle zoom for the current pane",
						},
						actionPair{
							action:      "Ctrl + b, x",
							description: "Kill the current pane",
						},
						actionPair{
							action:      "Ctrl + b, Space",
							description: "Toggle between pane layouts",
						},
					},
				},
				descriptiveResource{
					title: "tmux key bindings – windows",
					items: []resourceItem{
						actionPair{
							action:      "Ctrl + b, c",
							description: "Create a new window",
						},
						actionPair{
							action:      "Ctrl + b, d",
							description: "Detach from session",
						},
						actionPair{
							action:      "Ctrl + b, w",
							description: "List all tmux windows",
						},
						actionPair{
							action:      "Ctrl + b, &",
							description: "Kill the current window",
						},
						actionPair{
							action:      "Ctrl + b, n",
							description: "Navigate to the next window",
						},
						actionPair{
							action:      "Ctrl + b, p",
							description: "Navigate to the previous window",
						},
						actionPair{
							action:      "Ctrl + b, ,",
							description: "Rename the current window",
						},
						actionPair{
							action:      "Ctrl + b, 0...9",
							description: "Switch to window",
						},
					},
				},
				descriptiveResource{
					title: "tmux key bindings – utility",
					items: []resourceItem{
						actionPair{
							action:      "Ctrl + b, t",
							description: "Show the current time",
						},
						actionPair{
							action:      "Ctrl + b, ?",
							description: "Show all key bindings",
						},
					},
				},
				descriptiveResource{
					title: "tmux commands",
					items: []resourceItem{
						actionPair{
							action:      "Ctrl + b, :  or  tmux <command>",
							description: "Run a tmux command",
						},
						actionPair{
							action:      ":set status (on|off)",
							description: "Toggle status bar",
						},
						actionPair{
							action:      ":set status-position (top|bottom)",
							description: "Set the status bar position",
						},
						actionPair{
							action:      `:set status-style "bg=<color>,fg=<color>`,
							description: "Set status bar colors",
						},
					},
				},
			},
		},
	}
}
