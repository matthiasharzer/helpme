package resource

import (
	"strings"

	"github.com/fatih/color"
)

type Resource interface {
	Content() string
}

type DescriptiveResource struct {
	Title string
	Items []Item
}

func (dr DescriptiveResource) Content() string {
	titleColor := color.New(color.FgGreen).Add(color.Bold)
	var lines []string

	lines = append(lines, titleColor.Sprint(dr.Title))

	for _, kb := range dr.Items {
		lines = append(lines, kb.Text())
	}
	return strings.Join(lines, "\n")
}

type NestedResource struct {
	Items []Resource
}

func (nr NestedResource) Content() string {
	var lines []string
	for _, item := range nr.Items {
		lines = append(lines, item.Content())
	}
	return strings.Join(lines, "\n\n")
}
