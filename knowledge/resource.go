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

type actionPair struct {
	action      string
	description string
}

func (k actionPair) Text() string {
	key := color.New(color.FgMagenta).Add(color.Bold)
	return fmt.Sprintf("%s – %s", key.Sprint(k.action), k.description)
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
