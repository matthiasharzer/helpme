package tmux

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/matthiasharzer/helpme/knowledge/resource"
)

var tmuxPrefix string

func generateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func getTmuxPrefix() string {
	if tmuxPrefix != "" {
		return tmuxPrefix
	}

	tmuxLabel := fmt.Sprintf("tmux_temp_check_%s", generateRandomString(8))
	tmuxCmd := fmt.Sprintf("tmux -L %s -f ~/.tmux.conf start-server \\; show-options -gv prefix \\; kill-server", tmuxLabel)
	cmd := exec.Command("sh", "-c", tmuxCmd)
	output, err := cmd.Output()
	if err != nil {
		tmuxPrefix = "Ctrl + b"
	} else {
		tmuxPrefix = string(output)
		tmuxPrefix = strings.TrimSpace(tmuxPrefix)
		if strings.HasPrefix(tmuxPrefix, "C-") {
			tmuxPrefix = strings.Replace(tmuxPrefix, "C-", "Ctrl + ", 1)
		}
	}
	return tmuxPrefix
}

type PrefixResource struct {
	Action      string
	Description string
}

func (r PrefixResource) Text() string {
	prefix := getTmuxPrefix()

	key := resource.ActionPairItem{
		Action:      fmt.Sprintf("%s, %s", prefix, r.Action),
		Description: r.Description,
	}
	return key.Text()
}

type PrefixInfoItem struct {
	BeforePrefix string
}

func (p PrefixInfoItem) Text() string {
	prefix := getTmuxPrefix()
	prefixColor := color.New(color.FgMagenta).Add(color.Bold)
	return fmt.Sprintf("%s %s", p.BeforePrefix, prefixColor.Sprint(prefix))
}
