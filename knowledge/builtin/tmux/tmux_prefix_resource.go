package tmux

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/matthiasharzer/helpme/knowledge/resource"
)

var tmuxPrefix string

func getTmuxPrefix() string {
	if tmuxPrefix != "" {
		return tmuxPrefix
	}

	tmuxCmd := "tmux -L temp_check -f ~/.tmux.conf start-server \\; show-options -gv prefix \\; kill-server"
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
