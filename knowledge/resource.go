package knowledge

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Resource interface {
	Content() string
}

type resourceItem interface {
	Text() string
}

type keyBind struct {
	keys        string
	description string
}

func (k keyBind) Text() string {
	key := color.New(color.FgMagenta).Add(color.Bold)
	return fmt.Sprintf("%s – %s", key.Sprint(k.keys), k.description)
}

type descriptiveResource struct {
	title string
	items []resourceItem
}

func (dr descriptiveResource) Content() string {
	titleColor := color.New(color.FgGreen).Add(color.Bold)
	var lines []string

	lines = append(lines, titleColor.Sprint(dr.title))

	for _, kb := range dr.items {
		lines = append(lines, kb.Text())
	}
	return strings.Join(lines, "\n")
}

type nestedResource struct {
	items []Resource
}

func (nr nestedResource) Content() string {
	var lines []string
	for _, item := range nr.items {
		lines = append(lines, item.Content())
	}
	return strings.Join(lines, "\n\n")
}

func LoadResources() map[string]Resource {
	return map[string]Resource{
		"tmux": nestedResource{
			items: []Resource{
				descriptiveResource{
					title: "tmux key bindings – panes",
					items: []resourceItem{
						keyBind{
							keys:        "Ctrl + b, \"",
							description: "Split the current pane horizontally",
						},
						keyBind{
							keys:        "Ctrl + b, %",
							description: "Split the current pane vertically",
						},
						keyBind{
							keys:        "Ctrl + b, o",
							description: "Navigate to the next pane",
						},
						keyBind{
							keys:        "Ctrl + b, ;",
							description: "Close the current pane",
						},
						keyBind{
							keys:        "Ctrl + b, z",
							description: "Toggle zoom for the current pane",
						},
						keyBind{
							keys:        "Ctrl + b, x",
							description: "Kill the current pane",
						},
						keyBind{
							keys:        "Ctrl + b, Space",
							description: "Toggle between the last two panes",
						},
					},
				},
				descriptiveResource{
					title: "tmux key bindings – windows",
					items: []resourceItem{
						keyBind{
							keys:        "Ctrl + b, c",
							description: "Create a new windows",
						},
						keyBind{
							keys:        "Ctrl + b, d",
							description: "Detach from session",
						},
						keyBind{
							keys:        "Ctrl + b, w",
							description: "List all tmux windows",
						},
						keyBind{
							keys:        "Ctrl + b, &",
							description: "Kill the current window",
						},
						keyBind{
							keys:        "Ctrl + b, n",
							description: "Navigate to the next window",
						},
						keyBind{
							keys:        "Ctrl + b, p",
							description: "Navigate to the previous window",
						},
						keyBind{
							keys:        "Ctrl + b, ,",
							description: "Rename the current window",
						},
					},
				},
			},
		},
	}
}
