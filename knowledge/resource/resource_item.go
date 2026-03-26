package resource

import (
	"fmt"

	"github.com/fatih/color"
)

type Item interface {
	Text() string
}

type Text struct {
	TextContent string
}

func (t Text) Text() string {
	return t.TextContent
}

type ActionPairItem struct {
	Action      string
	Description string
}

func (i ActionPairItem) Text() string {
	key := color.New(color.FgMagenta).Add(color.Bold)
	return fmt.Sprintf("%s – %s", key.Sprint(i.Action), i.Description)
}
